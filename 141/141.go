package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快、慢指针求解回环问题
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next
	slow := head
	for slow != fast {
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		slow = slow.Next
	}
	return true
}

func test(head *ListNode, expect bool) {
	actual := hasCycle(head)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(nil, false)

	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	test(head, false)

	head.Next.Next.Next.Next = head
	test(head, true)
}
