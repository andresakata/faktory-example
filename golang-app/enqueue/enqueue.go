package main

import (
	"time"

	faktory "github.com/contribsys/faktory/client"
)

func main() {
	client, _ := faktory.Open()
	defer client.Close()

	for i := 1; i <= 6; i++ {
		job := faktory.NewJob("AnotherJob", i)
		job.Queue = "golang-app"
		client.Push(job)
		time.Sleep(10 * time.Second)
	}

	// scheduled job
	delay := 30 * time.Second
	futureTimeString := time.Now().Add(delay).UTC().Format(time.RFC3339)
	retryMaxAttempts := 2

	job := faktory.NewJob("AnotherJob", 7)
	job.At = futureTimeString
	job.Retry = &retryMaxAttempts
	job.Queue = "golang-app"

	client.Push(job)
}
