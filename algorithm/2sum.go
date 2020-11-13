package main
/* 
有序数组，给定一个数，求出数组中是否存在这两个元素
*/

func main() {
	a1 := []byte{1,2,3,4,5,6,7,8,9}
	print(IsExisted(18,a1))
}

func IsExisted(n byte,a []byte) bool {
	head := 0
	tail := len(a)-1
	println("start",head,tail)
	for {
		if head >= tail {
			return false
		}
		println(a[head],a[tail])
		if a[head] + a[tail] == n {
			return true
		}else if a[head] + a[tail] > n {
			tail = tail-1
		}else{
			head =head +1
		}
	}
	return false
}
