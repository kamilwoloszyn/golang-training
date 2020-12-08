package withcancel

import (
	"context"
	"fmt"
	"time"
)

func Run() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(301 * time.Millisecond)
		cancel()
		fmt.Println("ok")
	}()

	//After specific time make a decision
	select {
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Moing on")
	case <-ctx.Done():
		fmt.Println("work complete")
	}

}
