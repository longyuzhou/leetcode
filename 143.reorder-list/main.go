package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
// You may not modify the values in the list's nodes, only nodes itself may be changed.
// Example 1:
//   Given 1->2->3->4, reorder it to 1->4->2->3.
// Example 2:
//   Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
//
// Solution:
//   将链表分成两半，后一部分翻转
//     1->2->3->4->5 => 1->2->3
//                      5->4
//   然后将这两个链表合并为一个
//     1->2->3 => 1—>   2->   3
//     5->4    =>    5->   4->
//
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1, p2 := head, reverse(slow.Next)
	slow.Next = nil
	for p1 != nil && p2 != nil {
		next1 := p1.Next
		next2 := p2.Next
		p1.Next = p2
		p2.Next = next1
		p1 = next1
		p2 = next2
	}
}

func reverse(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	head := node
	p := node.Next
	head.Next = nil
	for p != nil {
		next := p.Next
		p.Next = head
		head = p
		p = next
	}
	return head
}

func build(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	return head
}

func test(nums []int, expect []int) {
	head := build(nums)
	reorderList(head)
	p := head
	pass := true
	actual := []int{}
	for _, num := range expect {
		if p == nil || p.Val != num {
			pass = false
		}
		if p == nil {
			actual = append(actual, -1)
		} else {
			actual = append(actual, p.Val)
			p = p.Next
		}
	}
	if !pass {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test([]int{1}, []int{1})
	test([]int{1, 2}, []int{1, 2})
	test([]int{1, 2, 3}, []int{1, 3, 2})
	test([]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3})
}
