package main

import (
	"log"
)

func main() {
	const N = 1024
	var a [N]int
	x := append(a[:N-1:N], 9, 9)
	y := append(a[:N:N], 9)
	z := append(append(a[:N-1:N],9),9)
	log.Println(len(x), cap(x), len(y), cap(y),len(z), cap(z))
}