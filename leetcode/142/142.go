package main


/*
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

假设链表头节点到环起点距离为x，环起点到相遇点为y，继续沿环顺时针走，相遇点到起点距离为z。
走到相遇点存在公式：2(a(y+z)+x+y) = b(y+z)+x+y  a,b表示慢快指针在环里的圈数。
简化：x = (b-2a)(y+z)-y
由于快慢指针，从相遇点在环里跑，慢指针跑一圈，快指针跑两圈，所以b=2a，所以x=z
*/

func main() {
	h2 := &ListNode{2, &ListNode{0, &ListNode{3, nil}}}
	h4 := &ListNode{-4, h2}
	h2.Next.Next.Next = h4
	_ = &ListNode{
		3,
		&ListNode{
			22,
			h2,
		},
	}
	log.Println(detectCycle(h2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next.Next
	slow := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			for slow != nil {
				if head == slow {
					return slow
				}
				head = head.Next
				slow = slow.Next
			}
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
