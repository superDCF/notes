package main

import (
	"database/sql"
	"runtime"
	"log"
)

const s = "Go101.org"

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

type A struct {
	a int
}

func main() {
	println(a, b)
	var x = *new(*int)
	println(x == nil)
	var y = new(*A)
	// (*y).a = 2
	log.Println(*y)
	runtime.StartTrace()
	
}
