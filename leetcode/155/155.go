package main

import "log"

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。

一个巧妙的解法：当最小值在栈里面，那么下次再插入的时候，直接还是插入该最小值
*/

func main() {
	obj := Constructor()
	obj.Push(512)
	obj.Push(-1024)
	obj.Push(-1024)
	obj.Push(512)
	obj.Pop()
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.GetMin())
}

type MinStack struct {
	Stack    []int
	MinQueue []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.MinQueue) == 0 || this.MinQueue[len(this.MinQueue)-1] >= val {
		this.MinQueue = append(this.MinQueue, val)
	} else {
		lastMin := this.MinQueue[len(this.MinQueue)-1]
		if lastMin < val {
			this.MinQueue = append(this.MinQueue, lastMin)
		}
	}
	log.Println("push", this.Stack, this.MinQueue)
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MinQueue = this.MinQueue[:len(this.MinQueue)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.MinQueue) == 0 {
		return 0
	}
	return this.MinQueue[len(this.MinQueue)-1]
}
