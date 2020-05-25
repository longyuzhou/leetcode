package main

import "fmt"

// 问题：给定已排序的链表，如果元素出现次数超过1，则移除该元素

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	head = &ListNode{}
	prev := head
	for curr != nil {
		next := curr
		count := 0
		for next != nil && curr.Val == next.Val {
			next = next.Next
			count++
		}
		if count == 1 {
			prev.Next = curr
			prev = curr
		}
		curr = next
	}
	prev.Next = nil
	return head.Next
}

func toList(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func test(nums []int, expect []int) {
	var head *ListNode
	if len(nums) > 0 {
		head = &ListNode{Val: nums[0]}
		p := head
		for i := 1; i < len(nums); i++ {
			p.Next = &ListNode{Val: nums[i]}
			p = p.Next
		}
	}
	head = deleteDuplicates(head)
	actual := toList(head)
	pass := len(expect) == len(actual)
	if pass {
		for i, num := range expect {
			if num != actual[i] {
				pass = false
				break
			}
		}
	}
	if !pass {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test([]int{1, 2, 2}, []int{1})
	test([]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5})
}
