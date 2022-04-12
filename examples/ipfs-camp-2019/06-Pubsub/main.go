package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	libp2p "github.com/riteshRcH/go-edge-device-lib"
	"github.com/riteshRcH/go-edge-device-lib/core/host"
	"github.com/riteshRcH/go-edge-device-lib/core/network"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	"github.com/riteshRcH/go-edge-device-lib/core/routing"
	kaddht "github.com/riteshRcH/go-edge-device-lib/dht"
	disc "github.com/riteshRcH/go-edge-device-lib/discovery"
	tls "github.com/riteshRcH/go-edge-device-lib/libp2ptls"
	"github.com/riteshRcH/go-edge-device-lib/multiaddr"
	"github.com/riteshRcH/go-edge-device-lib/p2p/discovery/mdns"
	mplex "github.com/riteshRcH/go-edge-device-lib/peerstream_multiplex"
	"github.com/riteshRcH/go-edge-device-lib/tcp"
	ws "github.com/riteshRcH/go-edge-device-lib/websocket"
)

type discoveryNotifee struct {
	h   host.Host
	ctx context.Context
}

func (m *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	if m.h.Network().Connectedness(pi.ID) != network.Connected {
		fmt.Printf("Found %s!\n", pi.ID.ShortString())
		m.h.Connect(m.ctx, pi)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
	)

	security := libp2p.Security(tls.ID, tls.New)

	listenAddrs := libp2p.ListenAddrStrings(
		"/ip4/0.0.0.0/tcp/0",
		"/ip4/0.0.0.0/tcp/0/ws",
	)

	var dht *kaddht.IpfsDHT
	newDHT := func(h host.Host) (routing.PeerRouting, error) {
		var err error
		dht, err = kaddht.New(ctx, h)
		return dht, err
	}
	routing := libp2p.Routing(newDHT)

	host, err := libp2p.New(
		transports,
		listenAddrs,
		muxers,
		security,
		routing,
	)
	if err != nil {
		panic(err)
	}

	// TODO: Replace our stream handler with a pubsub instance, and a handler
	// to field incoming messages on our topic.
	host.SetStreamHandler(chatProtocol, chatHandler)

	for _, addr := range host.Addrs() {
		fmt.Println("Listening on", addr)
	}

	targetAddr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/63785/p2p/QmWjz6xb8v9K4KnYEwP5Yk75k5mMBCehzWFLCvvQpYxF3d")
	if err != nil {
		panic(err)
	}

	targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		panic(err)
	}

	err = host.Connect(ctx, *targetInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connecting to bootstrap: %s", err)
	} else {
		fmt.Println("Connected to", targetInfo.ID)
	}

	notifee := &discoveryNotifee{h: host, ctx: ctx}
	mdns := mdns.NewMdnsService(host, "", notifee)
	if err := mdns.Start(); err != nil {
		panic(err)
	}

	err = dht.Bootstrap(ctx)
	if err != nil {
		panic(err)
	}

	routingDiscovery := disc.NewRoutingDiscovery(dht)
	disc.Advertise(ctx, routingDiscovery, string(chatProtocol))
	peers, err := disc.FindPeers(ctx, routingDiscovery, string(chatProtocol))
	if err != nil {
		panic(err)
	}
	for _, peer := range peers {
		notifee.HandlePeerFound(peer)
	}

	donec := make(chan struct{}, 1)
	go chatInputLoop(ctx, host, donec)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

	select {
	case <-stop:
		host.Close()
		os.Exit(0)
	case <-donec:
		host.Close()
	}
}
