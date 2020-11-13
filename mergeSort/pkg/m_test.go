package merge

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Logf("%v %v", Prev, Next)
	log.Println(Prev, Next)
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Merge(Prev, Next)
		Merge(Next, Prev)
	}
}

/*
Running tool: /home/qing/.g/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkMerge)$

goos: linux
goarch: amd64
BenchmarkMerge-6   	   15846	     74396 ns/op	  327680 B/op	       2 allocs/op
PASS
ok  	_/home/qing/golearn/src/other/blogs/mergeSort/pkg	1.949s
*/
/*
Running tool: /home/qing/.g/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkMerge2)$

goos: linux
goarch: amd64
BenchmarkMerge2-6   	   13544	     82234 ns/op	  327680 B/op	       2 allocs/op
PASS
ok  	_/home/qing/golearn/src/other/blogs/mergeSort/pkg	2.013s
*/
func BenchmarkMerge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Merge2(Prev, Next)
		Merge2(Next, Prev)
	}
}
