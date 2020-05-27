package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree, return the zigzag level order traversal
// of its nodes' values. (ie, from left to right,
// then right to left for the next level and alternate between).
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//    3
//   / \
//  9  20
//    /  \
//   15   7
// return its zigzag level order traversal as:
// [
//   [3],
//   [20,9],
//   [15,7]
// ]
//
// Solution: 层序遍历的变种
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	stack := []*TreeNode{root}
	leftToRight := true
	for len(stack) > 0 {
		size := len(stack)
		level := []int{}
		for i := 0; i < size; i++ {
			node := stack[i]
			if node != nil {
				stack = append(stack, node.Left, node.Right)
				level = append(level, node.Val)
			}
		}
		if !leftToRight {
			i, j := 0, len(level)-1
			for i < j {
				level[i], level[j] = level[j], level[i]
				i++
				j--
			}
		}
		stack = stack[size:]
		if len(level) > 0 {
			res = append(res, level)
		}
		leftToRight = !leftToRight
	}
	return res
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(nil))
	fmt.Println(zigzagLevelOrder(root))
}
