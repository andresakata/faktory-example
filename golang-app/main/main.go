package main

import (
	"context"
	"errors"
	"log"
	"time"

	worker "github.com/contribsys/faktory_worker_go"
)

func someFunc(ctx context.Context, args ...interface{}) error {
	help := worker.HelperFor(ctx)
	log.Printf("Processing %v, jid %s\n", args[0], help.Jid())

	if time.Now().Second() < 30 {
		return errors.New("Ops! Something went wrong.")
	}

	return nil
}

func main() {
	mgr := worker.NewManager()

	// register job types and the function to execute them
	mgr.Register("AnotherJob", someFunc)

	// use up to N goroutines to execute jobs
	mgr.Concurrency = 20

	// pull jobs from these queues, in this order of precedence
	mgr.ProcessStrictPriorityQueues("golang-app")

	// alternatively you can use weights to avoid starvation
	//mgr.ProcessWeightedPriorityQueues(map[string]int{"critical":3, "default":2, "bulk":1})

	// Start processing jobs, this method does not return.
	mgr.Run()
}
