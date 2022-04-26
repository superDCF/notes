package main

import (
	"errors"
	"fmt"
)

type T interface{}

func main() {
	var (
		t  T
		p  *T
		s1 interface{} = t
		s2 interface{} = p
		// s3 interface
	)
	//当声明一个变量为interface类型的时候，这个变量是nil，然后给这个interface赋了一个有类型的nil，则该变量不是nil，即interface不等于nil
	// 指针类型的anyType，值等于nil
	fmt.Println(s1 == t, s1 == nil, t == nil, (interface{})(nil) == nil, s1)  // true false
	fmt.Println(s2 == p, s2 == nil, p == nil, (*interface{})(nil) == nil, s2) // true true
	fmt.Printf("%v %T %v %T\n", s1, s1, s2, s2)
	fmt.Printf("%v %T %v %T\n", t, t, p, p)
	var s3 interface{}
	fmt.Println(s3 == nil)
	s3 = interface{}(nil)
	fmt.Println(s3 == nil)

	// fmt.Println(foo())
	s := []int{1, 2, 3, 4}
	Append(s)
	fmt.Println(s)
	Add(s)
	fmt.Println(s)
	// fmt.Println(moveK([]int{5, 1, 2, 3}))
}

func foo() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()
	defer func(e error) {
		fmt.Println(e)
		e = errors.New("b")
	}(err)
	return errors.New("c")
}

func Append(s []int) {
	s = append(s, 5)
}

func Add(s []int) {
	for i := range s {
		s[i] = s[i] + 5
	}
}

func maxFit(rocks []int) int {
	if len(rocks) == 0 {
		return 0
	}
	curFit := 0          // 当前的利润
	curPrice := rocks[0] // 当前价格
	isHas := false       // 是否持有股票，因为最多只能有一手，所以bool值表达
	for i := 1; i < len(rocks); i++ {
		if rocks[i-1] > rocks[i] && isHas { // 卖掉
			curFit = curFit + rocks[i-1] - curPrice
			curPrice = rocks[i-1]
			isHas = false
		} else {
			curPrice = rocks[i-1]
			isHas = true // 买
		}
	}
	return curFit
}

func moveK(arr []int) int {
	if len(arr) < 2 {
		return -1 // -1表示没找到
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return i + 1
		}
	}
	return -1
}
