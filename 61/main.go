package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	size := 1
	p := head
	for p.Next != nil {
		p = p.Next
		size++
	}

	k = k % size
	tail := p
	p = head
	for i := 0; i < size-k-1; i++ {
		p = p.Next
	}
	tail.Next = head
	head = p.Next
	p.Next = nil
	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	head = rotateRight(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
