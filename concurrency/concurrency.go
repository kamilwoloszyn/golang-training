package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func Run() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Goroutines")

	go func() {
		for i := 0; i < 60; i++ {
			fmt.Printf("A: %d \n", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 61; i++ {
			fmt.Printf("B: %d \n", i)
		}
		wg.Done()
	}()

	fmt.Println("Waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("Terminating program")
}
