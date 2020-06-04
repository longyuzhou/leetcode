package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表的插入排序
func insertionSortList(head *ListNode) *ListNode {
	place := func(node *ListNode) {
		var prev *ListNode
		curr := head
		for curr != nil && curr.Val < node.Val {
			prev = curr
			curr = curr.Next
		}
		if prev != nil {
			prev.Next = node
		} else {
			head = node
		}
		node.Next = curr
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		if prev != nil && prev.Val > curr.Val {
			node := curr
			prev.Next = curr.Next
			curr = curr.Next
			place(node)
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return head
}

func build(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	return head
}

func test(nums []int) {
	head := insertionSortList(build(nums))
	p := head
	actual := []int{}
	for p != nil {
		actual = append(actual, p.Val)
		p = p.Next
	}
	sort.Ints(nums)
	pass := len(nums) == len(actual)
	if pass {
		for i := range nums {
			if nums[i] != actual[i] {
				pass = false
				break
			}
		}
	}
	if pass {
		fmt.Println("pass")
	} else {
		fmt.Printf("fail! expect %v, got %v\n", nums, actual)
	}
}

func main() {
	test([]int{})
	test([]int{1})
	test([]int{4, 2, 1, 3})
}
