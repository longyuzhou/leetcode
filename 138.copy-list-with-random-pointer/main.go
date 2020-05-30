package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 深度拷贝链表（Node.Random指向该链表中任意一节点或nil）
func copyRandomList(head *Node) *Node {
	// 在原链表的每一个节点后插入该节点的拷贝
	// 1 -> 2 -> 3... => 1 -> 1' -> 2 -> 2' -> 3 -> 3'...
	p := head
	for p != nil {
		cp := &Node{Val: p.Val}
		next := p.Next
		p.Next = cp
		cp.Next = next
		p = next
	}

	// 处理节点拷贝的Random
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// 拆分成两个链表
	p = head
	var nhead *Node
	var np *Node
	for p != nil {
		cp := p.Next
		p.Next = p.Next.Next
		if nhead == nil {
			nhead = cp
			np = cp
		} else {
			np.Next = cp
			np = np.Next
		}
		p = p.Next
	}

	return nhead
}

func test(head *Node) {
	nhead := copyRandomList(head)
	p := head
	np := nhead
	i := 0
	for p != nil {
		if p == np {
			fmt.Printf("fail!第%d个节点未拷贝\n", i)
			return
		}
		if p.Val != np.Val {
			fmt.Printf("fail!第%d个节点的Val不同\n", i)
			return
		}
		if (p.Random == nil) != (np.Random == nil) {
			fmt.Printf("fail!第%d个节点的Random拷贝错误\n", i)
			return
		}
		if p.Random != nil && np.Random != nil {
			if p.Random == np.Random {
				fmt.Printf("fail!第%d个节点的Random未拷贝\n", i)
				return
			}
			if p.Random.Val != np.Random.Val {
				fmt.Printf("fail!第%d个节点的Random拷贝错误\n", i)
				return
			}
		}
		p = p.Next
		np = np.Next
		i++
	}
	fmt.Println("pass")
}

func build(data [][]int) *Node {
	var head *Node
	var tail *Node
	for _, e := range data {
		node := &Node{Val: e[0]}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	nodeAt := func(idx int) *Node {
		i := 0
		p := head
		for p != nil {
			if idx == i {
				return p
			}
			i++
			p = p.Next
		}
		return nil
	}

	i := 0
	p := head
	for p != nil {
		if len(data[i]) > 1 {
			p.Random = nodeAt(data[i][1])
		}
		i++
		p = p.Next
	}

	return head
}

func main() {
	test(build([][]int{[]int{1, 1}, []int{2, 1}}))
	test(build([][]int{[]int{7}, []int{13, 0}, []int{11, 4}, []int{10, 2}, []int{1, 0}}))
}
