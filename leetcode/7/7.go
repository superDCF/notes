package main

import (
	"log"
	"math"
)

func main() {
	defer func() {
		err := recover()
		log.Println("err=%v", err)
	}()
	log.Printf("%d", math.Pow(2, 31)-1)
	log.Println(reverse(-1234567899))
	log.Println(reverse(-123456789))
	// log.Println(int32(-9987654321))
}

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// 123 -> 321
func reverse(x int) int {
	n := 0
	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}
	if int(int32(n)) == n {
		return n
	}
	return 0
}
