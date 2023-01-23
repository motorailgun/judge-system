package queue_client

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
	ctx, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()

	connection, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Fprintf(os.Stderr, "gRPC Dial failed: %v\n", err)
		return &pb.Job{}, err
	}
	defer connection.Close()

	ctx, cancel2 := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel2()

	client := pb.NewJudgeServiceClient(connection)
	return client.GetJob(ctx, &pb.Empty{})
}
