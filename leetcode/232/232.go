package main

import "log"

/*
232. 用栈实现队列
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。

思路：
1. 栈1用于插入数据，栈2用于删除数据，如果删除时，栈2为空，则把栈1的数据加载到栈2，并清空栈1
*/

func main() {

	// log.Println(middleTree2(t, 20))
	q := Constructor()
	q.Push(1)
	log.Println(q.Pop(), q.Empty())
}

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	stackLen := len(this.stack1)
	if stackLen == 0 && len(this.stack2) == 0 {
		return 0
	}
	if len(this.stack2) == 0 {
		this.stack2 = make([]int, stackLen)
		for i, v := range this.stack1 {
			this.stack2[stackLen-1-i] = v
		}
		this.stack1 = this.stack1[:0]
	}
	ret := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	stackLen := len(this.stack1)
	if stackLen == 0 && len(this.stack2) == 0 {
		return 0
	}
	if len(this.stack2) == 0 {
		this.stack2 = make([]int, stackLen)
		for i, v := range this.stack1 {
			this.stack2[stackLen-1-i] = v
		}
		this.stack1 = this.stack1[:0]
	}
	ret := this.stack2[len(this.stack2)-1]
	return ret
}

func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	}
	return false
}
