package main

import (
	"log"
	"time"
)

var m = map[int]int{}

func main() {
	start := time.Now()
	log.Println(Fib(30))
	end := time.Since(start)
	log.Println(end.Milliseconds())
	log.Println(end.Nanoseconds())

	start = time.Now()
	log.Println(Fib2(30))
	end = time.Since(start)
	log.Println(end.Milliseconds())
	log.Println(end.Nanoseconds())

}

func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	if v, ok := m[n]; ok {
		return v
	}
	r := Fib(n-1) + Fib(n-2)
	m[n] = r
	/*
		{
			0:0,
			1:1,
			2:1,
			3:2,
			4:4
		}
	*/
	return r
}

func Fib2(n int) int {
	if n <= 2 {
		return 1
	}
	return Fib2(n-1) + Fib2(n-2)
}

/* 
2020/08/06 21:42:49 832040
2020/08/06 21:42:49 0
2020/08/06 21:42:49 85048
2020/08/06 21:42:49 832040
2020/08/06 21:42:49 2
2020/08/06 21:42:49 2779411
*/
