package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeoutContextBeforeTaskIsFinished()
	finishTaskBeforeContextTimeout()
}

func runTask(ctx context.Context, taskTime time.Duration) {
	select {
	case <-time.After(taskTime):
		fmt.Println("Finished long running task.")
	case <-ctx.Done():
		fmt.Println("Timed out context before task is finished.")
	}
}

func timeoutContextBeforeTaskIsFinished() {
	fmt.Println("This example should time out task execution:")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // Always cancel() to avoid context leak

	runTask(ctx, 3*time.Second)

	fmt.Println()
}

func finishTaskBeforeContextTimeout() {
	fmt.Println("This example should finish task execution before context timeout:")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Always cancel() to avoid context leak

	runTask(ctx, time.Second)

	fmt.Println()
}
