package main

import (
	"log"
	"strconv"
)

/*
字符串转换整数 (atoi)
*/

func main() {
	a1 := '+'
	log.Println(a1)
	a2 := int8(0)
	log.Println(^a2)
	log.Println(^1)
	log.Println(^2)
	log.Println(1<<8 - 1)
	var a3 = int8(0b1111111)
	log.Println(a3)
	log.Println(atoi2("- 1"))
	log.Println(strconv.Atoi("- 1"))
	log.Println(2 << 31)
	log.Println(atoi("   +-9999 11"))
	log.Println(atoi("   +9999 11"))
	log.Println(atoi("   9999 11"))
	log.Println(atoi("   -9999 11"))
	log.Println(atoi("words and 987"))
	log.Println(atoi("9223372036854775808"))
	log.Println(atoi("-91283472332"))
	log.Println(atoi("+ 0 123"))

}

func atoi(s string) int {
	ret := int(0)
	carry := 0
	signed := 1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && carry == 0 {
			continue
		} else if s[i] == '-' && carry == 0 {
			signed = -1
			carry++
		} else if s[i] == '+' && carry == 0 {
			carry++
		} else if s[i] >= '0' && s[i] <= '9' {
			n := s[i] - '0'
			ret = ret*10 + int(n)
			if ret > 1<<31-1 {
				ret = 1 << 31
				break
			}
			carry++
		} else {
			break
		}
	}
	ret = signed * ret
	if ret > 1<<31-1 {
		return (1<<31 - 1)
	}
	if ret < -(1 << 31) {
		return -(1 << 31)
	}
	return int(ret)
}

func atoi2(s string) int {
	sLen := len(s)
	if 0 < sLen && sLen < 19 {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0
			}
		}

		n := 0
		for _, ch := range []byte(s) {
			log.Println("ch", ch)
			ch -= '0'
			if ch > 9 { //负数用补码表示，补码是除符号位外，正整数的原码取反+1
				log.Println("ch2", ch)
				return 0
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n
	}
	return 0
}
