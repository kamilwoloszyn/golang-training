package tasks

import (
	"sync"
)

type Worker interface {
	Work()
}

type Task struct {
	work chan Worker
	wg   sync.WaitGroup
}

func MakeTasks(maxGoRoutines int) *Task {
	t := Task{
		work: make(chan Worker),
	}

	t.wg.Add(maxGoRoutines)
	for i := 0; i < maxGoRoutines; i++ {
		go func() {
			for w := range t.work {
				w.Work()
			}
			t.wg.Done()
		}()
	}
	return &t
}

func (t *Task) Shutdown() {
	close(t.work)
	t.wg.Wait()
}

//it submits work to the pool

func (t *Task) Do(w Worker) {
	t.work <- w
}
