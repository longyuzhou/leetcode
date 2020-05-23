package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	traverse(root, []*TreeNode{}, func(path []*TreeNode) {
		size := len(path)
		if size == 0 {
			return
		}
		if node := path[size-1]; node.Left == nil && node.Right == nil {
			s := ""
			for i, node := range path {
				s += strconv.Itoa(node.Val)
				if i < size-1 {
					s += "->"
				}
			}
			paths = append(paths, s)
		}
	})
	return paths
}

func traverse(node *TreeNode, path []*TreeNode, f func(path []*TreeNode)) {
	if node == nil {
		return
	}

	path = append(path, node)
	f(path)
	traverse(node.Left, path, f)
	traverse(node.Right, path, f)
}

func main() {
	paths := binaryTreePaths(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	})
	for _, path := range paths {
		fmt.Println(path)
	}
}
