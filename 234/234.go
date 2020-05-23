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
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 慢指针步进1
	// 快指针步进2
	// 慢指针位于第n / 2 + 1个元素，快指针位于第n / 2 * 2 + 1个元素
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil { // 链表长度为奇数，慢指针继续步进1
		slow = slow.Next
	}

	p1 := head
	p2 := reverseList(slow) // 翻转链表，然后比较
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

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

func test(head *ListNode, expect bool) {
	actutal := isPalindrome(head)
	if actutal != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(nil, false)
	test(&ListNode{Val: 1}, true)
	test(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, false)
	test(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}, true)
	test(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}, true)
}
