package main

import (
	"context"
	"fmt"
	"time"
)

func say(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Worker Started")
	case <-ctx.Done():
		fmt.Println("Worker Stopped")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // ✅ ensures resources are cleaned up

	go say(ctx)

	time.Sleep(10 * time.Second) // simulate some work
	fmt.Println("Main is canceling the worker...")
	cancel() // cancel after 1 second

	time.Sleep(4 * time.Second) // wait to see output
	fmt.Println("Main exits")
}
