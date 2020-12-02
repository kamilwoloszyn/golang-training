package atomicfunctions

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64

func incrementCoutner() {
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
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
