package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree, flatten it to a linked list in-place.
// For example, given the following tree:
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// The flattened tree should look like:
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
//
func flatten(root *TreeNode) {
	do(root)
}

func do(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	left := node.Left
	right := node.Right

	end := node
	if left != nil {
		end.Left = nil
		end.Right = left
		end = do(left)
	}
	if right != nil {
		end.Left = nil
		end.Right = right
		end = do(right)
	}
	end.Left = nil
	end.Right = nil
	return end
}

func print(root *TreeNode) {
	p := root
	for p != nil {
		fmt.Println(p.Val)
		p = p.Right
	}
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten(tree)
	print(tree)
}
