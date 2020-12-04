package buffered

import (
	"fmt"
	"runtime"
)

type Number struct {
	n int
}

type Parts struct {
	from int
	to   int
}

func MakeParts(numberOfElements int, numberOfGoroutines int) []Parts {
	var rangeOfIters int = int(numberOfElements / numberOfGoroutines)
	var result []Parts = make([]Parts, numberOfGoroutines)
	var last int
	for i := 0; i < numberOfGoroutines; i++ {
		result[i].from = last
		if last+rangeOfIters > numberOfElements {
			result[i].to = numberOfElements
		} else {
			result[i].to = last + rangeOfIters
		}

		last = result[i].to
	}
	return result
}

func init() {
	runtime.GOMAXPROCS(2)
}
func Run() {
	numberOfElements := 500
	numbers := make([]Number, numberOfElements)
	goRoutines := 10
	var p []Parts = MakeParts(numberOfElements, goRoutines)
	var waitChannels int = 10
	sum := make(chan int, 10)
	for i := 0; i < len(numbers); i++ {
		numbers[i].n = i
	}

	for i := 0; i < goRoutines; i++ {
		go func(num []Number, p Parts) {
			var sumPart int
			for i := p.from; i < p.to; i++ {
				sumPart += num[i].n
			}
			sum <- sumPart
		}(numbers, p[i])
	}
	fmt.Println("Generating result ")
	for waitChannels > 0 {
		fmt.Printf("Got: %d\n", <-sum)
		waitChannels--
	}
	fmt.Print("Done")

}
