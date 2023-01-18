package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"service_queue/pb"
	"service_queue/service"
)

func main() {
	port := 5001
	listenPort, err := net.Listen("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	pb.RegisterJudgeServiceServer(server, service.NewJudgeService())
}
