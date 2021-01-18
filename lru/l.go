/*
LRU 算法
1. Get & Put 都是O(1)
2. 有一个固定的长度

1. Get O(1) 选择Map实现；
2. Get 移动位置到组的第一个位置，表示为最近使用；
2. Put O(1) ，放置一个元素，采用双向链表实现，符合快速插入、移动
综上：考虑Map+Link实现
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	l := NewLru(3)
	log.Printf("a=%v", l.Get("a"))
	l.Put("a", 1)
	log.Printf("a=%v", l.Get("a"))
	l.Put("b", 2)
	log.Printf("b=%v l=%+v", l.Get("b"), l)
}

type Lru struct {
	hash map[string]*Node
	list *DoubleList
	Len  int
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Key  string
	Val  int
	Pre  *Node
	Next *Node
}

func NewLru(n int) *Lru {
	if n < 1 {
		return nil
	}
	l := &Lru{
		Len:  n,
		list: new(DoubleList),
		hash: make(map[string]*Node, n),
	}
	return l
}

func (l *Lru) Get(key string) int {
	v, ok := l.hash[key]
	if !ok {
		return -1
	}
	if l.list.size != 1 {
		l.list.addRecently(key, v)
	}
	return v.Val
}

func (l *Lru) Put(key string, val int) {
	node := &Node{
		Key: key,
		Val: val,
	}
	if l.list.size == l.Len {
		latestKey := l.list.deleteLatest()
		delete(l.hash, latestKey)
		l.list.Put(node)
	} else {
		l.hash[key] = node
		l.list.Put(node)
	}
}

func (l *Lru) String() string {
	s := fmt.Sprintf("hash=%+v list=%+v", l.hash, l.list)
	return s
}

func (d *DoubleList) addRecently(key string, val *Node) {
	d.delete(val)
	d.Put(val)
}

func (d *DoubleList) delete(node *Node) {
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		if node.Pre == nil {
			next := node.Next
			next.Pre = nil
			d.head = next
		} else if node.Next == nil {
			pre := node.Pre
			pre.Next = nil
			d.tail = pre
		} else {
			node.Pre.Next = node.Next
			node.Next.Pre = node.Pre
		}
	}
}

func (d *DoubleList) deleteLatest() (key string) {
	if d.size == 0 {
		return key
	} else if d.size == 1 {
		key = d.head.Key
		d.head = nil
		d.tail = nil
	} else {
		key = d.tail.Key
		d.tail = d.tail.Pre
		d.tail.Next = nil
	}
	d.size--
	return key
}

func (d *DoubleList) Put(node *Node) {
	if d.size == 0 {
		d.head = node
		d.tail = node
	} else {
		head := d.head
		head.Pre = node
		node.Next = head
		d.head = node
	}
	d.size++
}

func (d *DoubleList) String() string {
	s := ""
	for node := d.head; node != nil; {
		s += fmt.Sprintf("%+v", node)
		node = node.Next
	}
	return s
}
