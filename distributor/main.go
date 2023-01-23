package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"distributor/client"
	"distributor/judge"
)

const polling_interval int = 2
const additional_wait_on_failure int = 15

func main() {
	queue_addr := os.Getenv("QUEUE_ADDR")
	queue_port := os.Getenv("QUEUE_PORT")
	worker := client.NewJobGetter(queue_addr, queue_port)

	for {
		job, err := worker.GetJob()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get job from queue: %v", err)
			time.Sleep(time.Second * time.Duration(additional_wait_on_failure))
		}

		judger := judge.NewJudger()
		_, err = judger.Judge(job)

		time.Sleep(
			time.Second*time.Duration(polling_interval) + time.Millisecond*time.Duration(rand.Intn(200)),
		)
	}
}
