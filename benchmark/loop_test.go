package benchmark

import (
	"math/rand"
	"testing"
)

const size = 1000

func MakeArray() [size][size]int {
	var testArr [size][size]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			testArr[j][i] = rand.Int() % 30
		}
	}
	return testArr
}
func BenchmarkIterHorizontalOrder(b *testing.B) {
	var testArr [size][size]int = MakeArray()
	IterVerticalOrder(testArr)
}

func BenchmarkIterVerticalOrder(b *testing.B) {
	var testArr [size][size]int = MakeArray()
	IterVerticalOrder(testArr)
}

func BenchmarkModifyVerticalOrder(b *testing.B) {
	var testArr [size][size]int = MakeArray()
	ModifyVerticalOrder(testArr)
}

func BenchmarkModifyHorizontalOrder(b *testing.B) {
	var testArr [size][size]int = MakeArray()
	ModifyVerticalOrder(testArr)
}
