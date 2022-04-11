package util

import (
	"context"
	"fmt"
	"time"

	"github.com/riteshRcH/go-edge-device-lib/core/discovery"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"go.uber.org/zap"
)

var log, _ = zap.NewProduction()

// FindPeers is a utility function that synchronously collects peers from a Discoverer.
func FindPeers(ctx context.Context, d discovery.Discoverer, ns string, opts ...discovery.Option) ([]peer.AddrInfo, error) {
	var res []peer.AddrInfo

	ch, err := d.FindPeers(ctx, ns, opts...)
	if err != nil {
		return nil, err
	}

	for pi := range ch {
		res = append(res, pi)
	}

	return res, nil
}

// Advertise is a utility function that persistently advertises a service through an Advertiser.
func Advertise(ctx context.Context, a discovery.Advertiser, ns string, opts ...discovery.Option) {
	go func() {
		for {
			ttl, err := a.Advertise(ctx, ns, opts...)
			if err != nil {
				log.Debug(fmt.Sprintf("Error advertising %s: %s", ns, err.Error()))
				if ctx.Err() != nil {
					return
				}

				select {
				case <-time.After(2 * time.Minute):
					continue
				case <-ctx.Done():
					return
				}
			}

			wait := 7 * ttl / 8
			select {
			case <-time.After(wait):
			case <-ctx.Done():
				return
			}
		}
	}()
}
