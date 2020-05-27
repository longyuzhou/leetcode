package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 给定满二叉树, 将节点的Next指向同一层的右侧节点
//       1->nil
//      / \
//     /   \
//    /     \
//   2------>3->nil
//  / \     / \
// 4-->5-->6-->7->nil
//
// Solution: 如下算法支持任意类型二叉树(见117)
func connect(root *Node) *Node {
	prev := root
	prevIndex := 0

	// 获取下一个右节点
	get := func() *Node {
		for prev != nil {
			p := prev
			index := prevIndex
			if index%2 == 1 {
				prev = prev.Next
			}
			prevIndex++
			if index%2 == 0 && p.Left != nil { // 左节点
				return p.Left
			}
			if index%2 == 1 && p.Right != nil { // 右节点
				return p.Right
			}
		}
		return nil
	}

	for {
		// 找到p节点
		p := get()
		next := p
		if p == nil {
			break
		}
		// 处理当前层的节点
		for p != nil {
			p.Next = get()
			p = p.Next
		}
		// 进入下一层
		prev = next
		prevIndex = 0
	}
	return root
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	connect(root)
	p := root
	for p != nil {
		q := p
		for q != nil {
			fmt.Print(q.Val)
			if q.Next != nil {
				fmt.Print(", ")
			}
			q = q.Next
		}
		fmt.Print("\n")
		p = p.Left
	}
}
