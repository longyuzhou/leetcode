package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		idx := len(stack) - 1
		node := stack[idx]
		stack = stack[:idx]
		if node != nil {
			res = append(res, node.Val)
			stack = append(stack, node.Right, node.Left)
		}
	}
	return res
}

func test(root *TreeNode, expect []int) {
	actual := preorderTraversal(root)
	pass := len(expect) == len(actual)
	if pass {
		for i := range expect {
			if expect[i] != actual[i] {
				pass = false
				break
			}
		}
	}
	if !pass {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
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
			Val:   6,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
	}
	test(nil, []int{})
	test(root, []int{4, 2, 1, 3, 6, 5, 7})
}
