package main

import (
	"log"
	"reflect"
)

func main() {
	
	s := SortBySort([]T{T{5},T{4},T{3},T{2},T{1}},[]int{1,2,3,4,5})
	log.Printf("%+v",s)
	var t = []T{T{6},T{5},T{4},T{3},T{2},T{1}}
	r := Filter(t,func(e interface{}, i int, array []interface{}) interface{} {
		return e != nil
	})
	log.Printf("%+v",reflect.TypeOf(t).Kind(),r)
}

type T struct {
	F int
}

// 在un中某个字段按照给定的顺序排序
// 例如 [T{5},T{4},T{3},T{2},T{1}] 按照 [1,2,3,4,5]的顺序排序
func SortBySort(un []T, sort []int)[]T {
	m := make(map[int]int, len(un))
	for i, v := range un {
		m[v.F] = i
	}
	un2 := make([]T, len(sort))
	for i1, v := range sort {
		if i2, ok := m[v]; ok {
			un2[i1] = un[i2]
		}
	}
	return un2
}

type ItemHandler func(e interface{}, i int, array []interface{}) interface{}

func Filter(array interface{}, fn ItemHandler) []interface{} {
	if reflect.TypeOf(array).Kind() != reflect.Slice {
		log.Panic("no array")
	}
	// array.([]T)
	list := make([]interface{}, 0, 2)
	// for i, e := range array {
	// 	item := fn(e, i, array)
	// 	list = append(list, item)
	// }
	return list
}