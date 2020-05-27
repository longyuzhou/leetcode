package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given an integer n, generate all structurally unique
// BST's (binary search trees) that store values 1 ... n.
// Example:
//   Input: 3
//   Output:
//   [
//     [1,null,3,2],
//     [3,2,null,1],
//     [3,1,null,null,2],
//     [2,1,3],
//     [1,null,2,null,3]
//   ]
//   Explanation:
//   The above output corresponds to the 5 unique BST's shown below:
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
//
// Solution: 分治法
//   f(s, e) = f(s, i-1) + f(i+1, e)
//
func generateTrees(n int) []*TreeNode {
	var dfs func(s, e int) []*TreeNode
	dfs = func(s, e int) []*TreeNode {
		if s > e {
			return []*TreeNode{nil}
		}
		res := []*TreeNode{}
		for i := s; i <= e; i++ {
			leftNodes := dfs(s, i-1)
			rightNodes := dfs(i+1, e)
			for _, leftNode := range leftNodes {
				for _, rightNode := range rightNodes {
					node := &TreeNode{Val: i}
					node.Left = leftNode
					node.Right = rightNode
					res = append(res, node)
				}
			}
		}
		return res
	}
	if n < 1 {
		return []*TreeNode{}
	}
	return dfs(1, n)
}

func toList(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode, index int)
	dfs = func(node *TreeNode, index int) {
		if node == nil {
			return
		}
		if len(res) < index+1 {
			for i := len(res); i < index+1; i++ {
				res = append(res, -1)
			}
		}
		res[index] = node.Val
		dfs(node.Left, 2*index+1)
		dfs(node.Right, 2*index+2)
	}
	dfs(root, 0)
	return res
}

func main() {
	trees := generateTrees(3)
	for _, tree := range trees {
		fmt.Println(toList(tree))
	}
}
