package main

import (
	"math/rand"
	"log"
	"time"
	l "other/notes/shuffle/s"
)

func main() {
	runtime.GOMAXPROC(1)
	log.Printf("random=%v", random(1,2))
	log.Printf("shuffle=%v", shuffle([]int{0,1,2,3,4,5,6,7,8}))
	s:=l.Constructor([]int{1,2,3})
	log.Printf("s.Shuffle()=%v",s.Shuffle())
	nums := []int{1,2,3}
	n1 := nums
	var n2 =make([]int,3)
	copy(n2,nums)
	n1[0] = 4
	log.Printf("n1=%v n2=%v nums=%v",n1,n2,nums)
}

// 洗牌算法，就是把元素互相替换
func shuffle(arr []int)[]int {
	for i := 0; i < len(arr); i++ {
		index := random(i,len(arr))
		tem := arr[i]
		arr[i] = arr[index]
		arr[index] = tem
	}
	return arr
}

func random(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(b-a) + a
}
