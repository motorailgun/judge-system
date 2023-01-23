package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"service_queue/pb"
	"service_queue/service"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("QUEUE_PORT"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load port config for queue: %v\n", err)
		os.Exit(1)
	}

	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	pb.RegisterJudgeServiceServer(server, service.NewJudgeService())

	reflection.Register(server)
	server.Serve(listenPort)
}
