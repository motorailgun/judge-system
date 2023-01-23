package worker

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"worker/dist_worker"
	"worker/judge_manager"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lang := os.Getenv("JUDGE_LANGUAGE")
	port, err := strconv.Atoi(os.Getenv(lang))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load port config for worker: %v\n", err)
		os.Exit(1)
	}

	listen_port, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start listening on port %d: %v\n", port, err)
		os.Exit(2)
	}

	server := grpc.NewServer()
	dist_worker.RegisterLangServiceServer(server, judge_manager.NewJudgeManager(lang))

	reflection.Register(server)
	server.Serve(listen_port)
}
