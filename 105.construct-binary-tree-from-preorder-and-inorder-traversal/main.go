package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given preorder and inorder traversal of a tree, construct the binary tree.
// Note:
// You may assume that duplicates do not exist in the tree.
// For example, given
//   preorder = [3,9,20,15,7]
//   inorder = [9,3,15,20,7]
// Return the following binary tree:
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// Solution: preorder中靠前的是子树根节点; inorder中左侧是左子树, 右侧是右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexes := make(map[int]int) // inorder中元素和其下标
	for i, num := range inorder {
		indexes[num] = i
	}

	var dfs func(preStart, preEnd, inStart, inEnd int) *TreeNode
	dfs = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		node := &TreeNode{}
		node.Val = preorder[preStart]
		pos := indexes[node.Val]
		node.Left = dfs(preStart+1, preStart+pos-inStart, inStart, pos-1)
		node.Right = dfs(preStart+pos-inStart+1, preEnd, pos+1, inEnd)
		return node
	}
	return dfs(0, len(preorder)-1, 0, len(inorder)-1)
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
	root = buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	inorder(root)
}
