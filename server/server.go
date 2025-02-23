package server

import (
	"bufio"
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

	fmt.Printf("listening for messages on %s:%d\n", a.ListenAddr, a.ListenPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept connection: %w", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Printf("client connected, local addr: '%v', remote addr: '%v'", connection.LocalAddr(), connection.RemoteAddr())
	defer connection.Close()

	reader := bufio.NewReader(connection)
	for {
		msg, err := reader.ReadString('\n')
		if msg == "\n" {
			return
		}
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected")
				return
			}
			fmt.Printf("error reading string from connection %s:%v\n", connection.LocalAddr().String(), err)
			return
		}
		fmt.Printf("msg: %s \n", msg)
		_, err = connection.Write([]byte("Hi\n"))
		if err != nil {
			fmt.Printf("error writing message into connection: %v", err)
			return
		}
	}
}
