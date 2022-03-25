/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。
*/
package main

import (
	"log"
)

func main() {
	log.Println(hand(122))
	log.Println(hand(121))
	log.Println(itoa(-121))
	log.Println(itoa(8999999999999991210))
}

func hand(n int) bool {
	if n < 0 {
		return false
	}
	n0 := n
	n1 := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		n1 = n1*10 + d
	}
	if n1 == n0 {
		return true
	}
	return false
}

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func itoa(n int) string {
	neg := false
	if n < 0 {
		n = -n
		neg = true
	}
	s := ""
	for n != 0 {
		i := n % 10
		_s := digits[i]
		s = string(_s) + s
		n = n / 10
	}
	if neg {
		s = "-" + s
	}
	return s
}
