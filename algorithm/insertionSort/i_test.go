package main

import (
	"testing"
)

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort([]int{2, 4, 6, 5, 3, 1, 100, 1, 2, 3, 4, 5, 12, 32, 43, 53, 232, 23, 1212, 1213, 43})
	}
}


/* 
Running tool: /home/.g/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkSelectSort)$

goos: linux
goarch: amd64
BenchmarkSelectSort-6   	21183153	        56.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/home/golearn/src/other/algorithm/insertionSort	1.267s

Running tool: /home/.g/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkSelectSort)$

goos: linux
goarch: amd64
BenchmarkSelectSort-6   	 4136538	       287 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/home/golearn/src/other/algorithm/selectSort/s	1.485s

*/