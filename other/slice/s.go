package main

import (
	"log"
	"math"
	"unsafe"
)

func main() {
	var a = []int{0, 1, 2, 3, 4, 5}
	b := a[:3:5]
	log.Printf("b=%v", b, len(b), cap(b))
	// var aa = []int{0, 1, 2, 3, 4, 5}
	aa := make([]int, 13, 14)
	log.Println(unsafe.Alignof(aa))
	log.Println(unsafe.Sizeof(aa))
	log.Println(aa, len(aa), cap(aa), *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&aa)) + 8)), &aa)
	log.Printf("%p", &aa)
	aa = append(aa, 6, 7, 9)

	// cc := append(append(bb,10),11)
	log.Println(aa, len(aa), cap(aa), unsafe.Pointer(&aa), &aa)
	log.Printf("%p", &aa)
	// log.Println(bb, len(bb), cap(bb), unsafe.Pointer(&bb),&bb)
	// log.Printf("%p",&bb)
	index := 0
	log.Println(a[0:index], a[index+1:])
	// a = append(a[0:index], a[index+1:]...)
	copy(a[1:],a[2:])
	
	log.Println(a)
	count :=0
	log.Println(int(math.Ceil(float64(count) * 100 / float64(0))))
}
