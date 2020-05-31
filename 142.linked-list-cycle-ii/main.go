package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 检测链表是否存在回环，如果存在，则返回环路起始节点；反之，返回nil
// Solution: 快慢指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { // 相遇，存在回环
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func test(head *ListNode, expect *ListNode) {
	actual := detectCycle(head)
	if actual != expect {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	first := &ListNode{Val: 1}
	test(first, nil)

	second := &ListNode{Val: 2}
	first.Next = second
	second.Next = first
	test(first, first)

	third := &ListNode{Val: 3}
	fourth := &ListNode{Val: 4}
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = second
	test(first, second)
}
