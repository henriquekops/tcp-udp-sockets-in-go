package main

import (
	"./sockets"
	"flag"
	"os"
	"strings"
)

// Responsible for argument parsing
// Calls the 'socket' package for socket initialization
func main() {

	modePtr := flag.String("mode", "server", "Start mode - {client, server}")
	networkPtr := flag.String("network", "tcp", "Transport layer - {tcp, udp}")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()
	sockets.GetSocket(strings.ToLower(*modePtr), strings.ToLower(*networkPtr))
}