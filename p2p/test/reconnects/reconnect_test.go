package reconnect

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"testing"
	"time"

	bhost "github.com/riteshRcH/go-edge-device-lib/p2p/host/basic"
	"go.uber.org/zap"

	"github.com/riteshRcH/go-edge-device-lib/core/host"
	"github.com/riteshRcH/go-edge-device-lib/core/network"
	"github.com/riteshRcH/go-edge-device-lib/core/protocol"

	swarmt "github.com/riteshRcH/go-edge-device-lib/swarm/testing"

	u "github.com/riteshRcH/go-edge-device-lib/ipfs-util"
	"github.com/stretchr/testify/require"
)

var log, _ = zap.NewProduction()

func EchoStreamHandler(stream network.Stream) {
	c := stream.Conn()
	log.Debug(fmt.Sprintf("%s echoing %s", c.LocalPeer(), c.RemotePeer()))
	go func() {
		_, err := io.Copy(stream, stream)
		if err == nil {
			stream.Close()
		} else {
			stream.Reset()
		}
	}()
}

type sendChans struct {
	send   chan struct{}
	sent   chan struct{}
	read   chan struct{}
	close_ chan struct{}
	closed chan struct{}
}

func newSendChans() sendChans {
	return sendChans{
		send:   make(chan struct{}),
		sent:   make(chan struct{}),
		read:   make(chan struct{}),
		close_: make(chan struct{}),
		closed: make(chan struct{}),
	}
}

func newSender() (chan sendChans, func(s network.Stream)) {
	scc := make(chan sendChans)
	return scc, func(s network.Stream) {
		sc := newSendChans()
		scc <- sc

		defer func() {
			s.Close()
			sc.closed <- struct{}{}
		}()

		buf := make([]byte, 65536)
		buf2 := make([]byte, 65536)
		u.NewTimeSeededRand().Read(buf)

		for {
			select {
			case <-sc.close_:
				return
			case <-sc.send:
			}

			// send a randomly sized subchunk
			from := rand.Intn(len(buf) / 2)
			to := rand.Intn(len(buf) / 2)
			sendbuf := buf[from : from+to]

			log.Debug(fmt.Sprintf("sender sending %d bytes", len(sendbuf)))
			n, err := s.Write(sendbuf)
			if err != nil {
				log.Debug(fmt.Sprintln("sender error. exiting:", err))
				return
			}

			log.Debug(fmt.Sprintf("sender wrote %d bytes", n))
			sc.sent <- struct{}{}

			if n, err = io.ReadFull(s, buf2[:len(sendbuf)]); err != nil {
				log.Debug(fmt.Sprintln("sender error. failed to read:", err))
				return
			}

			log.Debug(fmt.Sprintf("sender read %d bytes", n))
			sc.read <- struct{}{}
		}
	}
}

// TestReconnect tests whether hosts are able to disconnect and reconnect.
func TestReconnect2(t *testing.T) {
	// TCP RST handling is flaky in OSX, see https://github.com/golang/go/issues/50254.
	// We can avoid this by using QUIC in this test.
	h1, err := bhost.NewHost(swarmt.GenSwarm(t, swarmt.OptDisableTCP), nil)
	require.NoError(t, err)
	h2, err := bhost.NewHost(swarmt.GenSwarm(t, swarmt.OptDisableTCP), nil)
	require.NoError(t, err)
	hosts := []host.Host{h1, h2}

	h1.SetStreamHandler(protocol.TestingID, EchoStreamHandler)
	h2.SetStreamHandler(protocol.TestingID, EchoStreamHandler)

	rounds := 8
	if testing.Short() {
		rounds = 4
	}
	for i := 0; i < rounds; i++ {
		log.Debug(fmt.Sprintf("TestReconnect: %d/%d\n", i, rounds))
		subtestConnSendDisc(t, hosts)
	}
}

// TestReconnect tests whether hosts are able to disconnect and reconnect.
func TestReconnect5(t *testing.T) {
	const num = 5
	hosts := make([]host.Host, 0, num)
	for i := 0; i < num; i++ {
		// TCP RST handling is flaky in OSX, see https://github.com/golang/go/issues/50254.
		// We can avoid this by using QUIC in this test.
		h, err := bhost.NewHost(swarmt.GenSwarm(t, swarmt.OptDisableTCP), nil)
		require.NoError(t, err)
		h.SetStreamHandler(protocol.TestingID, EchoStreamHandler)
		hosts = append(hosts, h)
	}

	rounds := 4
	if testing.Short() {
		rounds = 2
	}
	for i := 0; i < rounds; i++ {
		log.Debug(fmt.Sprintf("TestReconnect: %d/%d\n", i, rounds))
		subtestConnSendDisc(t, hosts)
	}
}

func subtestConnSendDisc(t *testing.T, hosts []host.Host) {
	ctx := context.Background()
	numStreams := 3 * len(hosts)
	numMsgs := 10

	if testing.Short() {
		numStreams = 5 * len(hosts)
		numMsgs = 4
	}

	ss, sF := newSender()

	for _, h1 := range hosts {
		for _, h2 := range hosts {
			if h1.ID() >= h2.ID() {
				continue
			}

			h2pi := h2.Peerstore().PeerInfo(h2.ID())
			log.Debug(fmt.Sprintf("dialing %s", h2pi.Addrs))
			if err := h1.Connect(ctx, h2pi); err != nil {
				t.Fatal("Failed to connect:", err)
			}
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numStreams; i++ {
		h1 := hosts[i%len(hosts)]
		h2 := hosts[(i+1)%len(hosts)]
		s, err := h1.NewStream(context.Background(), h2.ID(), protocol.TestingID)
		if err != nil {
			t.Error(err)
		}

		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			go sF(s)
			log.Debug(fmt.Sprintf("getting handle %d", j))
			sc := <-ss // wait to get handle.
			log.Debug(fmt.Sprintf("spawning worker %d", j))

			for k := 0; k < numMsgs; k++ {
				sc.send <- struct{}{}
				<-sc.sent
				log.Debug(fmt.Sprintf("%d sent %d", j, k))
				<-sc.read
				log.Debug(fmt.Sprintf("%d read %d", j, k))
			}
			sc.close_ <- struct{}{}
			<-sc.closed
			log.Debug(fmt.Sprintf("closed %d", j))
		}(i)
	}
	wg.Wait()

	for i, h1 := range hosts {
		log.Debug(fmt.Sprintf("host %d has %d conns", i, len(h1.Network().Conns())))
	}

	for _, h1 := range hosts {
		// close connection
		cs := h1.Network().Conns()
		for _, c := range cs {
			if c.LocalPeer() > c.RemotePeer() {
				continue
			}
			log.Debug(fmt.Sprintf("closing: %s", c))
			c.Close()
		}
	}

	<-time.After(20 * time.Millisecond)

	for i, h := range hosts {
		if len(h.Network().Conns()) > 0 {
			t.Fatalf("host %d %s has %d conns! not zero.", i, h.ID(), len(h.Network().Conns()))
		}
	}
}
