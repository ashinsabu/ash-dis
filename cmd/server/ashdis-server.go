package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ashinsabu/ash-dis/server"
)

func main() {
	listenPort := flag.Int("port", 6370, "Port to run the Ashdis server on")
	listenAddr := flag.String("addr", "localhost", "Address to run the Ashdis server on")
	flag.Parse()

	fmt.Printf("starting ashdis server on %s:%d\n", *listenAddr, *listenPort)
	ashidisServer := server.NewAshdisServer(server.AshdisServerOpts{
		ListenPort: *listenPort,
		ListenAddr: *listenAddr,
	})
	err := ashidisServer.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
