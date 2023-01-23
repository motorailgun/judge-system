package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"distributor/judge"
	"distributor/queue_client"
)

const polling_interval int = 2
const additional_wait_on_failure int = 15

func main() {
	queue_addr := os.Getenv("QUEUE_ADDR")
	queue_port := os.Getenv("QUEUE_PORT")
	worker := queue_client.NewJobGetter(queue_addr, queue_port)

	for {
		job, err := worker.GetJob()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get job from queue: %v", err)
			time.Sleep(time.Second * time.Duration(additional_wait_on_failure))
		}

		if job.ValidJob {
			judger := judge.NewJudger()
			_, err = judger.Judge(job.Submission)
		}

		time.Sleep(
			time.Second*time.Duration(polling_interval) + time.Millisecond*time.Duration(rand.Intn(200)),
		)
	}
}
