package main

import "log"

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

func main() {
	h8 := &ListNode{
		2, &ListNode{4, nil},
	}
	h1 := &ListNode{3, h8}
	h2 := &ListNode{1, &ListNode{9, &ListNode{1, h8}}}
	log.Println(getIntersectionNode(h1, h2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha0 := headA
	haLen := 0
	for ha0 != nil {
		haLen++
		ha0 = ha0.Next
	}
	log.Println("a len", haLen)
	hb0 := headB
	hbLen := 0
	for hb0 != nil {
		hbLen++
		hb0 = hb0.Next
	}
	log.Println("a len", hbLen)
	var ha1, hb1 *ListNode
	var dis int
	if haLen > hbLen {
		dis = haLen - hbLen
		ha1 = headA
		hb1 = headB
	} else {
		dis = hbLen - haLen
		hb1 = headA
		ha1 = headB
	}
	// if dis
	for ha1 != nil {
		ha1 = ha1.Next
		dis--
		if dis == 0 {
			for ha1 != nil && hb1 != nil {
				if ha1 == hb1 {
					return ha1
				}
				ha1 = ha1.Next
				hb1 = hb1.Next
			}
			return nil
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha := headA
	hb := headB
	for ha != nil || hb != nil {
		if ha != nil && ha == hb {
			return ha
		}
		if ha == nil && hb == nil {
			return nil
		}
		ha = ha.Next
		hb = hb.Next
		if ha == nil {
			ha = headB
		}
		if hb == nil {
			hb = headA
		}

	}
	return nil
}
