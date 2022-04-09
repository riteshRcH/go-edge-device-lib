package stream

import (
	"io"
	"time"

	"github.com/riteshRcH/go-edge-device-lib/core/mux"
)

type Stream interface {
	io.Reader
	io.Writer
	io.Closer

	CloseWrite() error
	CloseRead() error
	Reset() error

	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
}

type stream struct {
	mux.MuxedStream
}

//lint:ignore SA1019 // This needs to build with older versions.
func WrapStream(str mux.MuxedStream) *stream {
	return &stream{MuxedStream: str}
}
