/*
反转字符串
[a,b,c,d] => [d,c,b,a]
*/

package main

import "log"

func main() {
	log.Println(reverse([]string{"a", "b", "c", "d", "e"}))
	log.Println([]byte("ab c"), ' ', 'a', "a")
}

func reverse(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	tail := len(arr) - 1
	for head := 0; head < len(arr)/2; head++ {
		tmp := arr[tail]
		arr[tail] = arr[head]
		arr[head] = tmp
		tail--
	}
	return arr
}
