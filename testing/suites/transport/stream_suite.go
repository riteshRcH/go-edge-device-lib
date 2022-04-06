package ttransport

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	crand "crypto/rand"
	mrand "math/rand"

	"github.com/riteshRcH/go-edge-device-lib/core/network"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/transport"
	"github.com/riteshRcH/go-edge-device-lib/testing/race"

	ma "github.com/riteshRcH/go-edge-device-lib/multiaddr"
)

// VerboseDebugging can be set to true to enable verbose debug logging in the
// stream stress tests.
var VerboseDebugging = false

var randomness []byte

var StressTestTimeout = 1 * time.Minute

func init() {
	// read 1MB of randomness
	randomness = make([]byte, 1<<20)
	if _, err := crand.Read(randomness); err != nil {
		panic(err)
	}

	if timeout := os.Getenv("TEST_STRESS_TIMEOUT_MS"); timeout != "" {
		if v, err := strconv.ParseInt(timeout, 10, 32); err == nil {
			StressTestTimeout = time.Duration(v) * time.Millisecond
		}
	}
}

type Options struct {
	ConnNum   int
	StreamNum int
	MsgNum    int
	MsgMin    int
	MsgMax    int
}

func fullClose(t *testing.T, s network.MuxedStream) {
	if err := s.CloseWrite(); err != nil {
		t.Error(err)
		s.Reset()
		return
	}
	b, err := ioutil.ReadAll(s)
	if err != nil {
		t.Error(err)
	}
	if len(b) != 0 {
		t.Error("expected to be done reading")
	}
	if err := s.Close(); err != nil {
		t.Error(err)
	}
}

func randBuf(size int) []byte {
	n := len(randomness) - size
	if size < 1 {
		panic(fmt.Errorf("requested too large buffer (%d). max is %d", size, len(randomness)))
	}

	start := mrand.Intn(n)
	return randomness[start : start+size]
}

func debugLog(t *testing.T, s string, args ...interface{}) {
	if VerboseDebugging {
		t.Logf(s, args...)
	}
}

func echoStream(t *testing.T, s network.MuxedStream) {
	// echo everything
	var err error
	if VerboseDebugging {
		t.Logf("accepted stream")
		_, err = io.Copy(&logWriter{t, s}, s)
		t.Log("closing stream")
	} else {
		_, err = io.Copy(s, s) // echo everything
	}
	if err != nil {
		t.Error(err)
	}
}

type logWriter struct {
	t *testing.T
	W io.Writer
}

func (lw *logWriter) Write(buf []byte) (int, error) {
	lw.t.Logf("logwriter: writing %d bytes", len(buf))
	return lw.W.Write(buf)
}

func echo(t *testing.T, c transport.CapableConn) {
	var wg sync.WaitGroup
	defer wg.Wait()
	for {
		str, err := c.AcceptStream()
		if err != nil {
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer str.Close()
			echoStream(t, str)
		}()
	}
}

func serve(t *testing.T, l transport.Listener) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for {
		c, err := l.Accept()
		if err != nil {
			return
		}
		defer c.Close()

		wg.Add(1)
		debugLog(t, "accepted connection")
		go func() {
			defer wg.Done()
			echo(t, c)
		}()
	}
}

func SubtestStress(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID, opt Options) {
	msgsize := 1 << 11

	rateLimitN := 5000 // max of 5k funcs, because -race has 8k max.
	rateLimitChan := make(chan struct{}, rateLimitN)
	for i := 0; i < rateLimitN; i++ {
		rateLimitChan <- struct{}{}
	}

	rateLimit := func(f func()) {
		<-rateLimitChan
		f()
		rateLimitChan <- struct{}{}
	}

	writeStream := func(s network.MuxedStream, bufs chan<- []byte) {
		debugLog(t, "writeStream %p, %d MsgNum", s, opt.MsgNum)

		for i := 0; i < opt.MsgNum; i++ {
			buf := randBuf(msgsize)
			bufs <- buf
			debugLog(t, "%p writing %d bytes (message %d/%d #%x)", s, len(buf), i, opt.MsgNum, buf[:3])
			if _, err := s.Write(buf); err != nil {
				t.Errorf("s.Write(buf): %s", err)
				continue
			}
		}
	}

	readStream := func(s network.MuxedStream, bufs <-chan []byte) {
		debugLog(t, "readStream %p, %d MsgNum", s, opt.MsgNum)

		buf2 := make([]byte, msgsize)
		i := 0
		for buf1 := range bufs {
			i++
			debugLog(t, "%p reading %d bytes (message %d/%d #%x)", s, len(buf1), i-1, opt.MsgNum, buf1[:3])

			if _, err := io.ReadFull(s, buf2); err != nil {
				t.Errorf("io.ReadFull(s, buf2): %s", err)
				debugLog(t, "%p failed to read %d bytes (message %d/%d #%x)", s, len(buf1), i-1, opt.MsgNum, buf1[:3])
				continue
			}
			if !bytes.Equal(buf1, buf2) {
				t.Errorf("buffers not equal (%x != %x)", buf1[:3], buf2[:3])
			}
		}
	}

	openStreamAndRW := func(c network.MuxedConn) {
		debugLog(t, "openStreamAndRW %p, %d opt.MsgNum", c, opt.MsgNum)

		s, err := c.OpenStream(context.Background())
		if err != nil {
			t.Errorf("failed to create NewStream: %s", err)
			return
		}

		bufs := make(chan []byte, opt.MsgNum)
		go func() {
			writeStream(s, bufs)
			close(bufs)
		}()

		readStream(s, bufs)
		fullClose(t, s)
	}

	openConnAndRW := func() {
		debugLog(t, "openConnAndRW")

		var wg sync.WaitGroup
		defer wg.Wait()

		l, err := ta.Listen(maddr)
		if err != nil {
			t.Error(err)
			return
		}
		defer l.Close()

		wg.Add(1)
		go func() {
			defer wg.Done()
			serve(t, l)
		}()

		c, err := tb.Dial(context.Background(), l.Multiaddr(), peerA)
		if err != nil {
			t.Error(err)
			return
		}
		defer c.Close()

		// serve the outgoing conn, because some muxers assume
		// that we _always_ call serve. (this is an error?)
		wg.Add(1)
		go func() {
			defer wg.Done()
			debugLog(t, "serving connection")
			echo(t, c)
		}()

		var openWg sync.WaitGroup
		for i := 0; i < opt.StreamNum; i++ {
			openWg.Add(1)
			go rateLimit(func() {
				defer openWg.Done()
				openStreamAndRW(c)
			})
		}
		openWg.Wait()
	}

	debugLog(t, "openConnsAndRW, %d conns", opt.ConnNum)

	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < opt.ConnNum; i++ {
		wg.Add(1)
		go rateLimit(func() {
			defer wg.Done()
			openConnAndRW()
		})
	}
}

func SubtestStreamOpenStress(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	l, err := ta.Listen(maddr)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	count := 10000
	workers := 5

	if race.WithRace() {
		// the race detector can only deal with 8128 simultaneous goroutines, so let's make sure we don't go overboard.
		count = 1000
	}

	var (
		connA, connB transport.CapableConn
	)

	accepted := make(chan error, 1)
	go func() {
		var err error
		connA, err = l.Accept()
		accepted <- err
	}()
	connB, err = tb.Dial(context.Background(), l.Multiaddr(), peerA)
	if err != nil {
		t.Fatal(err)
	}
	err = <-accepted
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if connA != nil {
			connA.Close()
		}
		if connB != nil {
			connB.Close()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < workers; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					s, err := connA.OpenStream(context.Background())
					if err != nil {
						t.Error(err)
						return
					}
					wg.Add(1)
					go func() {
						defer wg.Done()
						fullClose(t, s)
					}()
				}
			}()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count*workers; i++ {
			str, err := connB.AcceptStream()
			if err != nil {
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				fullClose(t, str)
			}()
		}
	}()

	timeout := time.After(StressTestTimeout)
	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-timeout:
		t.Fatal("timed out receiving streams")
	case <-done:
	}
}

func SubtestStreamReset(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	var wg sync.WaitGroup
	defer wg.Wait()

	l, err := ta.Listen(maddr)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()

		muxa, err := l.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer muxa.Close()

		s, err := muxa.OpenStream(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
		defer s.Close()

		// Some transports won't open the stream until we write. That's
		// fine.
		_, _ = s.Write([]byte("foo"))

		time.Sleep(time.Millisecond * 50)

		_, err = s.Write([]byte("bar"))
		if err == nil {
			t.Error("should have failed to write")
		}

	}()

	muxb, err := tb.Dial(context.Background(), l.Multiaddr(), peerA)
	if err != nil {
		t.Fatal(err)
	}
	defer muxb.Close()

	str, err := muxb.AcceptStream()
	if err != nil {
		t.Error(err)
		return
	}
	str.Reset()
}

func SubtestStress1Conn1Stream1Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	SubtestStress(t, ta, tb, maddr, peerA, Options{
		ConnNum:   1,
		StreamNum: 1,
		MsgNum:    1,
		MsgMax:    100,
		MsgMin:    100,
	})
}

func SubtestStress1Conn1Stream100Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	SubtestStress(t, ta, tb, maddr, peerA, Options{
		ConnNum:   1,
		StreamNum: 1,
		MsgNum:    100,
		MsgMax:    100,
		MsgMin:    100,
	})
}

func SubtestStress1Conn100Stream100Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	SubtestStress(t, ta, tb, maddr, peerA, Options{
		ConnNum:   1,
		StreamNum: 100,
		MsgNum:    100,
		MsgMax:    100,
		MsgMin:    100,
	})
}

func SubtestStress50Conn10Stream50Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	SubtestStress(t, ta, tb, maddr, peerA, Options{
		ConnNum:   50,
		StreamNum: 10,
		MsgNum:    50,
		MsgMax:    100,
		MsgMin:    100,
	})
}

func SubtestStress1Conn1000Stream10Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	SubtestStress(t, ta, tb, maddr, peerA, Options{
		ConnNum:   1,
		StreamNum: 1000,
		MsgNum:    10,
		MsgMax:    100,
		MsgMin:    100,
	})
}

func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	SubtestStress(t, ta, tb, maddr, peerA, Options{
		ConnNum:   1,
		StreamNum: 100,
		MsgNum:    100,
		MsgMax:    10000,
		MsgMin:    1000,
	})
}
