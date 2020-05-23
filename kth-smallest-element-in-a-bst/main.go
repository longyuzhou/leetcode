package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	idx := 0
	node := root
	stack := []*TreeNode{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		size := len(stack)
		node = stack[size-1]
		idx++
		if idx == k {
			return node.Val
		}
		stack = stack[:size-1]

		node = node.Right
	}
	return -1
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	actual := kthSmallest(root, 1)
	fmt.Println(actual)
}
