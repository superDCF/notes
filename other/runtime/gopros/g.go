package main

import (
	"log"
	"runtime"
)

func main() {
	var a = new(*int)
	var b = 2
	*a= (&b)
	log.Printf("a=%v %v",*a,**a)
	log.Println(runtime.GOMAXPROCS(1))
	// var ch = make(chan int)
	// for {
	// 	go func(){
	// 		ch <- 1
	// 		fmt.Print("0")
	// 	}()
	// 	fmt.Print("1")

	// }
	

}