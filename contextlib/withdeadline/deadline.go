package withdeadline

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// to prevent block goroutine
	channel := make(chan struct{}, 1)

	go func() {
		time.Sleep(152 * time.Millisecond)
		channel <- struct{}{}
	}()

	select {
	case <-channel:
		fmt.Println("Done!")
	case <-ctx.Done():
		fmt.Println("Deadline exceed. Cancelling ..")
	}

}
