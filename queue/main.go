package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	port := 5001
	listenPort, err := net.Listen("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	server := grpc.NewServer()

}
