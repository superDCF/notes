package main

import (
	"log"

	. "./l"
)

func main() {
	l := NewList()
	l.Add(1)
	l.Add(3)
	l.Add(5)
	l.Add(4)
	l.Add(2)
	log.Printf("%s", l)
	l.Head = l.Reserve(l.Head)
	log.Printf("%s", l)
	l.Head = l.DeleteNode(l.Head, 5)
	log.Printf("%s", l)
}
