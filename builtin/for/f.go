package main

import "fmt"

/* 
for range 只要是取循环的地址就会有问题，因为每次都会用同一个地址存储循环出来的值，拷贝给该地址
*/

func main() {
	a := 1
	b := 2
	c := 3 
	d := 4
	e := 5
	slice := []*int{&a, &b, &c, &d, &e}
	myMap := make(map[int]**int)

	for index, value := range slice {
		myMap[index] = &value
	}
	fmt.Println("=====new map=====", myMap)

	slice2 := []int{0, 1, 2, 3}
	myMap2 := make(map[int]*int)
	for index, value := range slice2 {
		myMap2[index] = &value
	}
	fmt.Println("=====new map=====", myMap2)
	// prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
