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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	done := make(chan bool)
	go func() {
		traverse(root, func(balanced bool) {
			if !balanced {
				done <- false
			}
		})
		done <- true
	}()
	return <-done
}

func traverse(node *TreeNode, f func(balanced bool)) int {
	if node == nil {
		return 0
	}

	left := traverse(node.Left, f)
	right := traverse(node.Right, f)
	if abs(left-right) > 1 {
		f(false)
	}
	return 1 + max(left, right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func tree(vals []string, idx int) *TreeNode {
	if idx > len(vals)-1 {
		return nil
	}

	if vals[idx] == "" {
		return nil
	}

	val, _ := strconv.Atoi(vals[idx])
	node := &TreeNode{}
	node.Val = val
	node.Left = tree(vals, 2*idx+1)
	node.Right = tree(vals, 2*idx+2)
	return node
}

func test(root *TreeNode, expect bool) {
	actual := isBalanced(root)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(tree([]string{}, 0), true)
	test(tree([]string{"1", "", "2", "", "", "", "3"}, 0), false)
	test(tree([]string{"3", "9", "20", "", "", "15", "7"}, 0), true)
	test(tree([]string{"1", "2", "2", "3", "3", "", "", "4", "4"}, 0), false)
	test(tree([]string{"1", "2", "2", "3", "3", "3", "3", "4", "4", "4", "4", "4", "4", "", "", "5", "5"}, 0), true)
}
