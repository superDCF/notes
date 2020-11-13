package list

import (
	"fmt"
)

type Node struct {
	V int
	N *Node
}

type List struct {
	Head *Node
}

func NewList() *List {
	return &List{}
}

// Add 由小到大
func (l *List) Add(v int) {
	if l == nil {
		return
	}
	n := &Node{
		V: v,
	}
	if l.Head == nil {
		l.Head = n
		return
	}
	item := l.Head
	// item.N != nil 才能找到最后一个，哨兵机制
	for item != nil {
		// log.Println(item, v)
		if item.V > v {
			n := &Node{
				V: item.V,
				N: item.N,
			}
			item.V = v
			item.N = n
			return
		}

		if item.N == nil {
			break
		}
		item = item.N
	}
	item.N = n
	return
}

// reserve list 从某一个节点开始翻转
func (l *List) Reserve(n *Node) *Node {
	if n == nil || n.N == nil {
		return n
	}
	h := l.Reserve(n.N)
	n.N.N = n
	n.N = nil
	return h
}

func (l *List) String() string {
	s := ""
	for i := l.Head; i != nil; i = i.N {
		s += fmt.Sprintf("%v", i)
	}
	return s
}

func (l *List) DeleteNode(head *Node, val int) *Node {
	if head == nil {
		return nil
	}
	if head.V == val {
		return head.N
	}

	newHead := &Node{0, head}
	for head.N != nil {
		if head.N.V == val {
			head.N = head.N.N
		}
		head = head.N
	}
	return newHead.N
}
