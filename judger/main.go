package main

import (
	"net/http"
	"os"
)

func main() {
	queue_addr := os.Getenv("QUEUE_ADDR")
	queue_port := os.Getenv("QUEUE_PORT")

	response, err := http.Get(queue_addr + queue_port)

}
