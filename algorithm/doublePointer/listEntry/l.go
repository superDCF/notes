/*
链表找环，并找环的入口
*/

package main

import (
	"context"
	"log"
	"sync"
)

func main() {
	n6 := &node{6, nil}
	n3 := &node{3, &node{4, &node{5, n6}}}
	n6.next = n3
	list := &node{0, &node{1, &node{2, n3}}}
	list2 := &node{-6, &node{-5, &node{-4, &node{-3, &node{-2, &node{-1, list}}}}}}
	log.Println(entry(list))
	log.Println(entry(list2))
	context.Background()
}

type node struct {
	val  int
	next *node
}

func entry(root *node) (int, int) {
	fast, slow := root, root
	index := 0
	for fast != nil && fast.next != nil {
		// log.Println(slow.val, fast.val)
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			// n := slowIndex / (fastIndex - slowIndex)
			start := root
			for slow.next != nil {
				if slow == start {
					entry := slow
					ringLen := 1
					for slow.next != entry {
						slow = slow.next
						ringLen++
					}
					return index, ringLen
				}
				slow = slow.next
				start = start.next
				index++
			}
		}
	}
	return -1, -1
}
