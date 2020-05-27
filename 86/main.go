package main

import (
	"fmt"

	utils "github.com/longyuzhou/leetcode"
)

func test(nums []int, x int, expect []int) {
	head := partition(utils.ListToLinkedList(nums), x)
	actual := utils.LinkedListToList(head)
	if !utils.SameList(expect, actual) {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

// 问题：给链表分区，小于指定值的排在前面，大于等于的排后面，同时保持原顺序不变
func partition(head *utils.ListNode, x int) *utils.ListNode {
	h1 := &utils.ListNode{}
	t1 := h1
	h2 := &utils.ListNode{}
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
