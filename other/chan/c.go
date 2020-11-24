package main

import (
	"log"
	"runtime"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	log.Println("c3",c)
	// close(c)
	c <- sum // send sum to c
	log.Println("c33",c)
	c <- sum+1
	log.Println("c4",c)
	close(c)
	log.Println("c2",c)
}

func main() {
	num := runtime.GOMAXPROCS(1)
	log.Println(num, runtime.NumCPU())
	s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int,3)
	// c = nil
	var c chan int
	log.Println("c1",c)
	c <- 1
	log.Println("c11",c)
	go sum(s[:len(s)/2], c)
	time.Sleep(time.Second)
	log.Println("c5",c)
	// close(c)
	// go sum(s[len(s)/2:], c)
	x, y := <-c // receive from c
	// var x1 int
	// var y1 bool
	// select {
	// case x1, y1 = <-c:
	// default:
	// }
	// log.Println(x, y,x1, y1)
	log.Println(x, y)
}
