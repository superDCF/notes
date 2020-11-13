/*
实现Redis的List
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	d := new(DoubleList)

	log.Println(d.Len)
	d.LPush(0)
	d.LPush(1)
	d.LPush(2)
	// d.LPush(3)
	// d.LPush(4)
	log.Printf("len=%v d=%+v", d.Len, d)
	log.Printf("key=0 d=%+v", d.LIndex(0))
	log.Printf("key=1 d=%+v", d.LIndex(1))
	log.Printf("key=-2 d=%+v", d.LPop())
	log.Printf("key=-1 d=%+v", d.LIndex(-1))
	log.Printf("len=%v d=%+v", d.Len, d)
}

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type DoubleList struct {
	head *Node
	tail *Node
	Len  int
}

// func NewDoubleList(n int) *DoubleList {
// 	d := &DoubleList{
// 		Len: n,
// 	}
// 	return d
// }

func (d *DoubleList) LIndex(i int) *Node {
	if d.Len < i {
		return nil
	}
	var node *Node

	if i < 0 { // 倒数第几个
		i = d.Len + i
	}

	// 从头开始查，还是从尾开始查
	if d.Len/2 > i { // 从头遍历查
		node = d.head
		for j := 1; j <= i && i > 0; j++ {
			node = node.Next
		}
	} else {
		node = d.tail
		for j := d.Len - 1 - 1; j >= i && i < d.Len-1; j-- {
			node = node.Prev
		}
	}
	return node
}

func (d *DoubleList) LPush(val int) {
	node := &Node{
		Val: val,
	}
	if d.Len == 0 {
		d.head = node
		d.tail = node

	} else {
		head := d.head
		head.Prev = node
		node.Next = head
		d.head = node
	}
	d.Len++
}

func (d *DoubleList) LPop() *Node {
	if d.Len == 0 {
		return nil
	}
	node := d.head
	next := d.head.Next
	next.Prev = nil
	d.head = next
	d.Len--
	return node
}

func (d *DoubleList) String() string {
	s := ""
	for node := d.head; node != nil; {
		s += fmt.Sprintf("%+v", node)
		node = node.Next
	}
	return s
}
