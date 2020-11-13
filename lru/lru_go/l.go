package main

import (
	"container/list"
	"fmt"
	"log"
)

func main() {
	l := NewLru(3)
	fmt.Printf("a=%v\n", l.Get("a"))
	l.Put("a", 1)
	fmt.Printf("a=%v\n", l.Get("a"))
	fmt.Printf("b=%v\n", l.Get("b"))
	l.Put("b", 2)
	fmt.Printf("b=%v\n", l.Get("b"))
	fmt.Printf("c=%v\n", l.Get("c"))
	l.Put("c", 3)
	fmt.Printf("c=%v\n", l.Get("c"))
	l.Put("d", 3)
	fmt.Printf("d=%v\n", l.Get("d"))
	// fmt.Printf("a=%v\n", l.Get("a"))
	fmt.Printf("l=%+v lruLen=%v listLen=%v\n", l, l.Len, l.list.Len())
}

type Lru struct {
	hash map[string]*list.Element
	list *list.List
	Len  int
}

func NewLru(n int) *Lru {
	return &Lru{
		hash: make(map[string]*list.Element),
		list: list.New(),
		Len:  n,
	}
}

func (l *Lru) String() string {
	s := fmt.Sprintf("hash=%+v list=%+v\n", l.hash, l.list)
	for n := l.list.Front(); n != nil; {
		s += fmt.Sprintf("ele=%+v", n)
		n = n.Next()
	}
	return s
}

func (l *Lru) Get(key string) int {
	v, ok := l.hash[key]
	if !ok {
		return -1
	}

	m, ok := v.Value.(map[string]int)
	if !ok {
		return -1
	}
	l.list.MoveToFront(v)
	return m[key]
}

func (l *Lru) Put(key string, val int) {
	node := &list.Element{
		Value: map[string]int{key: val},
	}
	if l.list.Len() == l.Len {
		last := l.list.Back()
		v := l.list.Remove(last)
		m, ok := v.(map[string]int)
		if !ok {
			return
		}
		for k := range m {
			delete(l.hash, k)
		}
	}
	l.hash[key] = node
	l.list.PushFront(node)
	log.Println("key=", key)
}
