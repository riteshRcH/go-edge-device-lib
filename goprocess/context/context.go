package goprocessctx

import (
	"context"

	goprocess "github.com/riteshRcH/go-edge-device-lib/goprocess"
)

// WithContext constructs and returns a Process that respects
// given context. It is the equivalent of:
//
//   func ProcessWithContext(ctx context.Context) goprocess.Process {
//     p := goprocess.WithParent(goprocess.Background())
//     CloseAfterContext(p, ctx)
//     return p
//   }
//
func WithContext(ctx context.Context) goprocess.Process {
	p := goprocess.WithParent(goprocess.Background())
	CloseAfterContext(p, ctx)
	return p
}

// WithContextAndTeardown is a helper function to set teardown at initiation
// of WithContext
func WithContextAndTeardown(ctx context.Context, tf goprocess.TeardownFunc) goprocess.Process {
	p := goprocess.WithTeardown(tf)
	CloseAfterContext(p, ctx)
	return p
}

// WaitForContext makes p WaitFor ctx. When Closing, p waits for
// ctx.Done(), before being Closed(). It is simply:
//
//   p.WaitFor(goprocess.WithContext(ctx))
//
func WaitForContext(ctx context.Context, p goprocess.Process) {
	p.WaitFor(WithContext(ctx))
}

// CloseAfterContext schedules the process to close after the given
// context is done. It is the equivalent of:
//
//   func CloseAfterContext(p goprocess.Process, ctx context.Context) {
//     go func() {
//       <-ctx.Done()
//       p.Close()
//     }()
//   }
//
func CloseAfterContext(p goprocess.Process, ctx context.Context) {
	if p == nil {
		panic("nil Process")
	}
	if ctx == nil {
		panic("nil Context")
	}

	// Avoid a goroutine for both context.Background() and goprocess.Background().
	if ctx.Done() == nil || p.Closed() == nil {
		return
	}

	go func() {
		select {
		case <-ctx.Done():
			p.Close()
		case <-p.Closed():
		}
	}()
}

// WithProcessClosing returns a context.Context derived from ctx that
// is cancelled as p is Closing (after: <-p.Closing()). It is simply:
//
//   func WithProcessClosing(ctx context.Context, p goprocess.Process) context.Context {
//     ctx, cancel := context.WithCancel(ctx)
//     go func() {
//       <-p.Closing()
//       cancel()
//     }()
//     return ctx
//   }
//
func WithProcessClosing(ctx context.Context, p goprocess.Process) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	p.AddChildNoWait(goprocess.WithTeardown(func() error {
		cancel()
		return nil
	}))
	return ctx
}

// WithProcessClosed returns a context.Context that is cancelled
// after Process p is Closed. It is the equivalent of:
//
//   func WithProcessClosed(ctx context.Context, p goprocess.Process) context.Context {
//     ctx, cancel := context.WithCancel(ctx)
//     go func() {
//       <-p.Closed()
//       cancel()
//     }()
//     return ctx
//   }
//
func WithProcessClosed(ctx context.Context, p goprocess.Process) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	p.AddChildNoWait(goprocess.WithTeardown(func() error {
		select {
		case <-p.Closed():
		case <-ctx.Done():
		}
		cancel()
		return nil
	}))
	return ctx
}
