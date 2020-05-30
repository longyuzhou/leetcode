package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

// Given a reference of a node in a connected undirected graph.
// Return a deep copy (clone) of the graph.
// Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.
//
// Solution: 动态规划
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	dp := make(map[int]*Node)

	var dfs func(node *Node)
	dfs = func(node *Node) {
		copy, ok := dp[node.Val]
		if ok {
			return
		}
		copy = &Node{Val: node.Val}
		dp[node.Val] = copy
		if node.Neighbors != nil {
			copy.Neighbors = []*Node{}
			for _, neighbor := range node.Neighbors {
				dfs(neighbor)
				copy.Neighbors = append(copy.Neighbors, dp[neighbor.Val])
			}
		}
	}

	dfs(node)
	return dp[node.Val]
}

func main() {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	one.Neighbors = []*Node{two, four}
	two.Neighbors = []*Node{three, one}
	three.Neighbors = []*Node{four, two}
	four.Neighbors = []*Node{one, three}
	fmt.Println(one)
	fmt.Println(cloneGraph(one))
}
