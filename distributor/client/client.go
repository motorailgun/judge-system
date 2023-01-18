package client

import (
	"context"
	"distributor/pb"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type JobGetter struct {
	server_address string
	server_port    string
}

func NewJobGetter(address string, port string) JobGetter {
	return JobGetter{
		server_address: address,
		server_port:    port,
	}
}

func (worker *JobGetter) GetJob() (*pb.Job, error) {
	address := worker.server_address + ":" + worker.server_port
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Fprintf(os.Stderr, "gRPC Dial failed: %v\n", err)
		return &pb.Job{}, err
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client := pb.NewJudgeServiceClient(connection)
	return client.GetJob(ctx, &pb.Empty{})
}
