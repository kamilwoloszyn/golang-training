//GOLANG BENCHMARK

package structures

import (
	"testing"
)

var dummy uint32

func BenchmarkIterOverLinkedList(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverLinkedList()
	}
	dummy = result
}

func BenchmarkIterOverMatrix(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverMatrix()
	}
	dummy = result
}

func BenchmarkIterOverMatrixRow(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverMatrixRow()
	}
	dummy = result
}

func BenchmarkIterOverMatrixColumn(b *testing.B) {
	var result uint32
	for i := 0; i < b.N; i++ {
		result = IterOverMatrixColumn()
	}
	dummy = result
}
