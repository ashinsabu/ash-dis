package server

import (
	"fmt"
	"net"
)

type AshdisServer struct {
	AshdisServerOpts
}

type AshdisServerOpts struct {
	ListenAddr string `envconfig:"ASHDIS_ADDR" default:"localhost"`
	ListenPort int    `envconfig:"ASHDIS_PORT" default:"6370"`
}

func NewAshdisServer(ashdisServerOpts AshdisServerOpts) *AshdisServer {
	return &AshdisServer{
		AshdisServerOpts: ashdisServerOpts,
	}
}

// test setup a basic tcp server that just listens for connections and responds Hello world
func (a *AshdisServer) Start() error {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(a.ListenAddr),
		Port: a.ListenPort,
	})
	if err != nil {
		return fmt.Errorf("failed to listen on %s:%d: %w", a.ListenAddr, a.ListenPort, err)
	}
	defer listener.Close()

	fmt.Printf("Listening for messages on %s:%d\n", a.ListenAddr, a.ListenPort)
	return nil
}
