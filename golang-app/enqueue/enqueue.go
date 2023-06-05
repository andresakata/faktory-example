package main

import (
	"time"

	faktory "github.com/contribsys/faktory/client"
)

func main() {
	client, _ := faktory.Open()
	defer client.Close()

	for i := 1; i <= 6; i++ {
		job := faktory.NewJob("SomeJob", i)
		client.Push(job)
		time.Sleep(10 * time.Second)
	}

	// scheduled job
	delay := 30 * time.Second
	futureTimeString := time.Now().Add(delay).UTC().Format(time.RFC3339)
	retryMaxAttempts := 2

	job := faktory.NewJob("SomeJob", 7)
	job.At = futureTimeString
	job.Retry = &retryMaxAttempts
	job.Queue = "default"

	client.Push(job)
}
