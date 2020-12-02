package mutexes

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

type Data struct {
	head    string
	content []string
}

func read(d *Data, wg *sync.WaitGroup) {
	rwMutex.RLock()
	fmt.Printf("Head: %s, Content: %v\n", d.head, d.content)
	rwMutex.RUnlock()
	wg.Done()
}

func write(d *Data, wg *sync.WaitGroup) {
	rwMutex.Lock()
	d.head = "Modified"
	d.content = append(d.content, "y")
	fmt.Printf("Append: %s, %v\n", d.head, d.content)
	rwMutex.Unlock()
	wg.Done()
}

func getRandomStr(length int) []string {
	rand.Seed(time.Now().UnixNano())
	var result []string
	for len(result) < length {
		if tmp := rand.Intn(0x7a); tmp > 0x61 {
			result = append(result, string(rune(tmp)))
		}
	}

	return result
}

func init() {
	runtime.GOMAXPROCS(2)
}

func Run() {

	var wg sync.WaitGroup
	var readRoutines = 10
	var writeRoutines = 1

	data := Data{
		"Helo",
		getRandomStr(15),
	}

	for i := 0; i < readRoutines; i++ {
		wg.Add(1)
		go read(&data, &wg)

	}

	for j := 0; j < writeRoutines; j++ {
		wg.Add(1)
		go write(&data, &wg)

	}

	wg.Wait()
	fmt.Println("Done")

}
