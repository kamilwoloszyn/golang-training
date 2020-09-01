//GOLANG BENCHMARK

package structures

import (
	"fmt"
	"testing"
)

func BenchmarkIterOverLinkedList(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverLinkedList()
	}
	fmt.Printf("\nComplete. Result: %d\n", result)
}

func BenchmarkIterOverMatrix(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverMatrix()
	}
	fmt.Printf("\nComplete. Result: %d\n", result)
}

func BenchmarkIterOverMatrixRow(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverMatrixRow()
	}
	fmt.Printf("\nComplete. Result: %d\n", result)
}

func BenchmarkIterOverMatrixColumn(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverMatrixColumn()
	}
	fmt.Printf("\nComplete. Result: %d\n", result)
}
