package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	head = nil
	for p != nil {
		tmp := p.Next
		p.Next = head
		head = p
		p = tmp
	}
	return head
}

func build(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	p := head
	for i := 1; i < len(nums); i++ {
		p.Next = &ListNode{Val: nums[i]}
		p = p.Next
	}
	return head
}

func traverse(head *ListNode) {
	p := head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

func main() {
	traverse(reverseList(build([]int{1, 2, 3, 4, 5})))
}
