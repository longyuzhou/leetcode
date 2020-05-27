package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a singly linked list where elements are sorted in ascending order,
// convert it to a height balanced BST.
//
// Solution: 使用快慢指针来找到中点
func sortedListToBST(head *ListNode) *TreeNode {
	return build(head, nil)
}

func build(head *ListNode, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}
	if head.Next == tail {
		return &TreeNode{Val: head.Val}
	}

	fast, slow := head, head
	for fast.Next != tail && fast.Next.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	node := &TreeNode{Val: slow.Val}
	node.Left = build(head, slow)
	node.Right = build(slow.Next, tail)
	return node
}

func buildLinkedList(nums []int, index int) *ListNode {
	var node *ListNode
	if len(nums) > index {
		node = &ListNode{}
		node.Val = nums[index]
		node.Next = buildLinkedList(nums, index+1)
	}
	return node
}

func inorder(root *TreeNode) []int {
	res := []int{}
	p := root
	stack := []*TreeNode{}
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		idx := len(stack) - 1
		p = stack[idx]
		stack = stack[:idx]
		res = append(res, p.Val)

		p = p.Right
	}
	return res
}

func test(nums []int) {
	linkedList := buildLinkedList(nums, 0)
	tree := sortedListToBST(linkedList)
	actual := inorder(tree)
	pass := len(nums) == len(actual)
	if pass {
		for i, num := range nums {
			if num != actual[i] {
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
	test([]int{-10, -3, 0, 5, 9})
}
