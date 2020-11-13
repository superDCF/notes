package main

import (
	"log"

	"./s"
)

func main() {
	log.Println(s.SelectSort([]int{2, 4, 6, 5, 3, 1, 100}))
}
