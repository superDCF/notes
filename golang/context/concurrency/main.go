package main

import "fmt"

// go version go1.16.1 darwin/amd64 
const (
	x = iota
	_
	y
	z = "zz"
	k 
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
