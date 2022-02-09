package main

import (
	"fmt"
	"sync"
)

type Student struct {
	Name string
	Id   int
}

func main() {
	s := make(map[string]Student)
	s["chenchao"] = Student{
		Name: "chenchao",
		Id:   111,
	}
	a := s["chenchao"]
	a.Id = 222
	aa := make([]int, 3)
	fmt.Println(cap(aa))

	b := make(map[int]int, 3)
	fmt.Println(len(b),b)
	sync.Map
}
