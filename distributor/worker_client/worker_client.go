package worker_client

import (
	"context"
	"distributor/dist_worker"
	"distributor/pb"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerClient struct {
	address string
	port    string
}

func NewWorkerClient(port string) WorkerClient {
	return WorkerClient{
		"localhost",
		port,
	}
}

func (client *WorkerClient) RunJudge(submission *pb.Submission) (*pb.JudgeResult, error) {
	address := client.address + ":" + client.port
	ctx, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Fprintf(os.Stderr, "worker connection failed: %v\n", err)
		return &pb.JudgeResult{
			ResultCode: pb.JudgeInfo_err,
			Submission: submission,
		}, err
	}
	defer conn.Close()

	ctx, cancel2 := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel2()

	cli := dist_worker.NewLangServiceClient(conn)
	return cli.RunJudge(ctx, submission)
}
