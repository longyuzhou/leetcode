package main

import (
	"fmt"

	utils "github.com/longyuzhou/leetcode"
)

// Reverse a linked list from position m to n. Do it in one-pass.
// Note: 1 ≤ m ≤ n ≤ length of list.
// Example:
//   Input: 1->2->3->4->5->NULL, m = 2, n = 4
//   Output: 1->4->3->2->5->NULL
//
func reverseBetween(head *utils.ListNode, m int, n int) *utils.ListNode {
	var prev *utils.ListNode
	var h *utils.ListNode
	var t *utils.ListNode
	i := 1
	p := head
	for p != nil {
		next := p.Next
		if i == m-1 {
			prev = p
		} else if i == m {
			h = p
			t = p
			t.Next = nil
		}
		if m < i && i <= n {
			p.Next = h
			h = p
		}
		if i == n {
			t.Next = next
			break
		}
		i++
		p = next
	}
	if prev != nil {
		prev.Next = h
		return head
	}
	return h
}

func test(nums []int, m int, n int, expect []int) {
	head := utils.ListToLinkedList(nums)
	head = reverseBetween(head, m, n)
	actual := utils.LinkedListToList(head)
	if utils.SameList(actual, expect) {
		fmt.Println("pass")
	} else {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	}
}

func main() {
	test([]int{}, 0, 0, []int{})
	test([]int{1, 2, 3, 4, 5}, 1, 3, []int{3, 2, 1, 4, 5})
	test([]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5})
	test([]int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1})
	test([]int{1, 2}, 1, 3, []int{2, 1})
	test([]int{1, 2}, 1, 1, []int{1, 2})
}
