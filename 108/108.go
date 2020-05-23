package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return impl(nums, 0, len(nums)-1)
}

func impl(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := int(uint(low+high) >> 1)
	node := TreeNode{
		Val:   nums[mid],
		Left:  impl(nums, low, mid-1),
		Right: impl(nums, mid+1, high),
	}
	return &node
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left)
	fmt.Println(node.Val)
	traverse(node.Right)
}

func test(nums []int) {
	root := sortedArrayToBST(nums)
	traverse(root)
}

func main() {
	test([]int{-10, -3, 0, 5, 9})
}
