package main

import "log"

func main() {
	bar()
}

func foo() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("foo recover")
	// 	}
	// }()

	panic("foo")
}

func bar() {

	defer func() { // defer cover can recove panic in call stack
		if err := recover(); err != nil {
			log.Println("bar")
		}
	}()
	// defer recover() // err case

	foo()
	if err := recover(); err != nil {
		log.Println("bar2")
	}
}
