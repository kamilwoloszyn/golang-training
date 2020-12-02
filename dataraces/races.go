package dataraces

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

func incrementCoutner() {
	for i := 0; i < 10; i++ {
		val := counter
		runtime.Gosched()
		val++
		counter = val
	}

}

func Run() {
	const grs = 2
	var wg sync.WaitGroup
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		go func() {
			incrementCoutner()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter: %d", counter)

}
