package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree, return the inorder traversal of its nodes' values.
// Example:
//   Input: [1,null,2,3]
//      1
//       \
//        2
//       /
//      3
//   Output: [1,3,2]
// Follow up: Recursive solution is trivial, could you do it iteratively?
//
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	p := root
	stack := []*TreeNode{}
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		size := len(stack) - 1
		p = stack[size]
		stack = stack[:size]
		res = append(res, p.Val)

		p = p.Right
	}

	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(root))
}
