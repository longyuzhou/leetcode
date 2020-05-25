package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func toList(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func toLinkedList(nums []int) *ListNode {
	var head *ListNode
	if len(nums) > 0 {
		head = &ListNode{Val: nums[0]}
		p := head
		for i := 1; i < len(nums); i++ {
			p.Next = &ListNode{Val: nums[i]}
			p = p.Next
		}
	}
	return head
}

func sameList(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i, num := range nums1 {
		if num != nums2[i] {
			return false
		}
	}
	return true
}

func test(nums []int, x int, expect []int) {
	head := partition(toLinkedList(nums), x)
	actual := toList(head)
	if !sameList(expect, actual) {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

// 问题：给链表分区，小于指定值的排在前面，大于等于的排后面，同时保持原顺序不变
func partition(head *ListNode, x int) *ListNode {
	h1 := &ListNode{}
	t1 := h1
	h2 := &ListNode{}
	t2 := h2

	p := head
	for p != nil {
		if p.Val < x {
			t1.Next = p
			t1 = p
		} else {
			t2.Next = p
			t2 = p
		}
		p = p.Next
	}

	t1.Next = h2.Next
	t2.Next = nil
	return h1.Next
}

func main() {
	test([]int{}, 3, []int{})
	test([]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5})
}
