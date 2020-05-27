package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given inorder and postorder traversal of a tree, construct the binary tree.
// 和105类似
func buildTree(inorder []int, postorder []int) *TreeNode {
	indexes := make(map[int]int) // inorder中元素和其下标
	for i, num := range inorder {
		indexes[num] = i
	}

	var dfs func(postStart, postEnd, inStart, inEnd int) *TreeNode
	dfs = func(postStart, postEnd, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		node := &TreeNode{}
		node.Val = postorder[postEnd]
		pos := indexes[node.Val]
		node.Left = dfs(postStart, postStart+pos-inStart-1, inStart, pos-1)
		node.Right = dfs(postStart+pos-inStart, postEnd-1, pos+1, inEnd)
		return node
	}
	return dfs(0, len(postorder)-1, 0, len(inorder)-1)
}

func inorder(root *TreeNode) {
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
		fmt.Println(p.Val)

		p = p.Right
	}
}

func main() {
	root := buildTree([]int{}, []int{})
	root = buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	inorder(root)
}
