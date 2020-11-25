package main

import (
	"log"
	"unsafe"
)

func main() {
	var a = []int{0, 1, 2, 3, 4, 5}
	b := a[:3:5]
	log.Printf("b=%v", b, len(b), cap(b))
	// var aa = []int{0, 1, 2, 3, 4, 5}
	aa := make([]int, 10, 10)
	log.Println(aa, len(aa), cap(aa), uintptr(unsafe.Pointer(&aa))+10,&aa)
	log.Printf("%p",&aa)
	aa = append(aa, 6, 7, 9)

	// cc := append(append(bb,10),11)
	log.Println(aa, len(aa), cap(aa), unsafe.Pointer(&aa),&aa)
	log.Printf("%p",&aa)
	// log.Println(bb, len(bb), cap(bb), unsafe.Pointer(&bb),&bb)
	// log.Printf("%p",&bb)
	// log.Println(cc, len(cc), cap(cc))
}
