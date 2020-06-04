package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Sort a linked list in O(n log n) time using constant space complexity.
// Solution: 使用快慢指针找到链表中点，将链表分割为两半，执行归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	p1, p2 := head, slow
	prev.Next = nil
	return merge(sortList(p1), sortList(p2))
}

func merge(p1, p2 *ListNode) *ListNode {
	temp := &ListNode{}
	curr := temp
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	}
	if p2 != nil {
		curr.Next = p2
	}
	return temp.Next
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
	head := sortList(build(nums))
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

func print(head *ListNode) {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	fmt.Println(nums)
}

func main() {
	test([]int{4, 3, 2, 1, 0})
}
