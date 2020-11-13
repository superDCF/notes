package s

import (
	"testing"
)

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort([]int{2, 4, 6, 5, 3, 1, 100, 1, 2, 3, 4, 5, 12, 32, 43, 53, 232, 23, 1212, 1213, 43})
	}
}
