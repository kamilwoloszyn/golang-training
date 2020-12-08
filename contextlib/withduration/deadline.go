package withduration

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	duration := 20 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	ch := make(chan string, 1)
	defer cancel()
	go func() {
		time.Sleep(15 * time.Millisecond)
		ch <- "Ok"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Cancelled")
	case <-ch:
		fmt.Println("Done")
	}
}
