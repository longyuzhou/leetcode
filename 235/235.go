package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 针对binary search tree的特定解法
func bstLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
	for node != nil {
		mid := node.Val
		if p.Val < mid && q.Val < mid {
			node = node.Left
		} else if p.Val > mid && q.Val > mid {
			node = node.Right
		} else {
			return node
		}
	}
	return nil
}

// 这是通用解法，不只适用于binary search tree
// 如下所有的index指代当用一维数组来表示树时，树节点在该数组中的下标
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lowestCommonAncestorIndex := -1
	var lowestCommonAncestor *TreeNode
	pIndex := -1
	qIndex := -1
	traverse(root, 0, func(index int, node *TreeNode) bool {
		if p == node {
			pIndex = index
		} else if q == node {
			qIndex = index
		}
		if lowestCommonAncestorIndex == -1 && pIndex != -1 && qIndex != -1 {
			lowestCommonAncestorIndex = getLowestCommonAncestorIndex(pIndex, qIndex)
		}
		if index == lowestCommonAncestorIndex {
			lowestCommonAncestor = node
			return false
		}
		return true
	})
	return lowestCommonAncestor
}

func getLowestCommonAncestorIndex(pIndex, qIndex int) int {
	pAncestors := getAncestorIndexes(pIndex)
	qAncestors := getAncestorIndexes(qIndex)
	i := len(pAncestors) - 1
	j := len(qAncestors) - 1
	for i > 0 && j > 0 {
		if pAncestors[i-1] != qAncestors[j-1] {
			break
		}
		i--
		j--
	}
	return pAncestors[i]
}

func getAncestorIndexes(index int) []int {
	indexes := []int{index}
	for index > 0 {
		indexes = append(indexes, (index-1)/2)
		index = (index - 1) / 2
	}
	return indexes
}

func traverse(node *TreeNode, index int, callback func(index int, node *TreeNode) bool) bool {
	if node == nil {
		return true
	}

	return traverse(node.Left, 2*index+1, callback) &&
		traverse(node.Right, 2*index+2, callback) &&
		callback(index, node)
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}
	fmt.Println(lowestCommonAncestor(nil, root.Left, root.Right))
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right))

	fmt.Println("==========")

	fmt.Println(bstLowestCommonAncestor(nil, root.Left, root.Right))
	fmt.Println(bstLowestCommonAncestor(root, root.Left, root.Right))
	fmt.Println(bstLowestCommonAncestor(root, root.Left, root.Left.Right))
}
