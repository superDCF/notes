package main

func main() {

}

// 判断单链表有环，并且找到环的入口

// 思路：利用快慢2倍指针判断
/*
1. 首先判断是否有环
2. 有环的情况下，快指针一定会追上慢指针，因为每次移动只是快慢指针之间的距离+1，当在环中，环的大小是有限的，一定会在慢指针的在环内未走完至多一周多（走几周，与环的奇偶有关）内相遇。
设：入口点距离链表头为a
相遇点距离入口点逆时针最小距离是b
环的剩余长度是c

当相遇时，则有：
fast = 2 * slow
fast = slow + n(b+c)
fast = a + b + n(b+c)
slow = a + b

=> slow = n(b+c)
a = n(b+c) - b = (n-1)b+nc = (n-1)b + (n-1)c + c = (n-1)(b+c) + c
// 从上得知，slow在相遇点，再往前走c距离，就到环入口
也可以这样理解：
在相遇点每次走 (n-1)(b+c) + c 都会走到入口处，所以当前位置，再往前走c步，就到入口处
所以再找一个指针，让他从头开始走a步，slow走c步，在入口处就会相遇
*/

type Node struct {
	Next *Node
}

func circleEntry(head *Node) *Node {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}else{
			return nil
		}
		if fast == slow { // 相遇了,总有一轮会相遇的
			fast = head
			for fast != slow { // 在入口处就会相遇
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}
