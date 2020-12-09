package tasks

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	content string
}

func (d Data) Work() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(d.content)
}

func Run() {
	var wg sync.WaitGroup
	d := Data{
		"sampleContent",
	}
	t := MakeTasks(10)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			t.Do(d)
			wg.Done()
		}()
	}

	wg.Wait()
	t.Shutdown()
	fmt.Println("Done")
}
