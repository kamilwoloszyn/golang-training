package unbuffered

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(5)
}

func player1(c chan int, wg *sync.WaitGroup) {
	if _, ok := <-c; ok {
		fmt.Println("Player1: Ball recived")
	}
	wg.Done()

}

func player2(c chan int, wg *sync.WaitGroup) {
	//	time.Sleep(3 * time.Second)
	if _, ok := <-c; ok {
		fmt.Println("Player2: Ball recived")
	}
	wg.Done()
}

func Run() {

	ball := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go player1(ball, &wg)
	go player2(ball, &wg)
	ball <- 1
	fmt.Println("Finished")
	close(ball)

}
