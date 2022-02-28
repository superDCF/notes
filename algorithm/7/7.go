package main

/*
整数反转
123=》321
超过32位，即越界为0
*/

func main() {

}

func rever(x int) int {
	n := 0
	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}
	if int(int32(n)) != n {
		return 0
	}
	return n
}
