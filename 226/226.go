package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(tmp)
	return root
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left)
	fmt.Println(node.Val)
	traverse(node.Right)
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
	invertTree(root)
	traverse(root)
}
