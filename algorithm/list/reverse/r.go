// 链表反转

package main

func main() {
	
}

type Node struct {
	Val int
	Next *Node
}

// Reverse 1->2->3->4->5->6
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	last := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}