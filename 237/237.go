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
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	head := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}}
	deleteNode(head.Next)
	p := head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
