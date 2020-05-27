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

// Given a binary tree and a sum,
// find all root-to-leaf paths where each path's sum equals the given sum.
// Note: A leaf is a node with no children.
// Example:
// Given the below binary tree and sum = 22,
//        5
//       / \
//      4   8
//     /   / \
//    11  13  4
//   /  \    / \
//  7    2  5   1
// Return:
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]
//
// Solution: dfs
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}

	var dfs func(node *TreeNode, total int, path []int)
	dfs = func(node *TreeNode, total int, path []int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && total+node.Val == sum {
			res = append(res, append([]int{}, path...))
		}
		dfs(node.Left, total+node.Val, path)
		dfs(node.Right, total+node.Val, path)
	}

	dfs(root, 0, []int{})
	return res
}

func build(nums []string, index int) *TreeNode {
	if index >= len(nums) || nums[index] == "" {
		return nil
	}
	node := &TreeNode{}
	node.Val, _ = strconv.Atoi(nums[index])
	node.Left = build(nums, 2*index+1)
	node.Right = build(nums, 2*index+2)
	return node
}

func main() {
	nums := []string{}
	tree := build(nums, 0)
	fmt.Println(pathSum(tree, 0))

	nums = []string{"5", "4", "8", "11", "", "13", "4", "7", "2", "", "", "", "", "5", "1"}
	tree = build(nums, 0)
	fmt.Println(pathSum(tree, 22))
}
