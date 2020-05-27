package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree, determine if it is a valid binary search tree (BST).
// Assume a BST is defined as follows:
//  1) The left subtree of a node contains only nodes with keys less than the node's key.
//  2) The right subtree of a node contains only nodes with keys greater than the node's key.
//  3) Both the left and right subtrees must also be binary search trees.
//
// Solution: 使用中序遍历，prev.Val >= p.Val则表示顺序错误
func isValidBST(root *TreeNode) bool {
	p := root
	stack := []*TreeNode{}
	var prev *TreeNode
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		idx := len(stack) - 1
		p = stack[idx]
		stack = stack[:idx]
		if prev != nil && prev.Val >= p.Val {
			return false
		}
		prev = p
		p = p.Right
	}
	return true
}
