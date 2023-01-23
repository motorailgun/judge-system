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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Fprintf(os.Stderr, "worker connection failed: %v\n", err)
		return &pb.JudgeResult{
			ResultCode: pb.JudgeInfo_err,
			Submission: submission,
		}, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	cli := dist_worker.NewLangServiceClient(conn)
	return cli.RunJudge(ctx, submission)
}
