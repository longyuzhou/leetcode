package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针处理相遇问题
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 先假设两链表长度相同
	p1 := headA
	p2 := headB
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1 == nil && p2 == nil {
		return nil
	}

	// 两链表不等长，则将较长链表的指针移到与另一链表相对应的起始位置
	if p1 != nil {
		p := headA
		for p1 != nil {
			p = p.Next
			p1 = p1.Next
		}
		p1 = p
		p2 = headB
	} else {
		p := headB
		for p2 != nil {
			p = p.Next
			p2 = p2.Next
		}
		p1 = headA
		p2 = p
	}
	return getIntersectionNode(p1, p2)
}

func main() {
	fmt.Println(getIntersectionNode(nil, nil))

	fmt.Println(getIntersectionNode(&ListNode{Val: 1}, &ListNode{Val: 1}))

	head := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: head,
		},
	}
	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: head,
			},
		},
	}
	fmt.Println(getIntersectionNode(headA, headB))
}
