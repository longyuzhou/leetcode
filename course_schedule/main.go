package main

import "fmt"

// There are a total of numCourses courses you have to take,
// labeled from 0 to numCourses-1.
//
// Some courses may have prerequisites,
// for example to take course 0 you have to first take course 1,
// which is expressed as a pair: [0,1]
//
// Given the total number of courses and a list of prerequisite pairs,
// is it possible for you to finish all courses?
//
// Example 1:
//   Input: numCourses = 2, prerequisites = [[1,0]]
//   Output: true
//   Explanation: There are a total of 2 courses to take.
//                To take course 1 you should have finished course 0.
//                So it is possible.
//
// Example 2:
//   Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
//   Output: false
//   Explanation: There are a total of 2 courses to take.
//                To take course 1 you should have finished course 0,
//                and to take course 0 you should
//                also have finished course 1. So it is impossible.
//
// Solution:
//   (1) 统计所有顶点的入边数量，如果没有入边为0的顶点，无解；反之，执行(2)
//   (2) 移除入边为0的顶点及其出边，然后执行(1)
// 例1:
//   2 -> 1, 1 -> 0, 0 -> 2 没有入边为0的顶点，无解
//
// 例2:
//   3 -> 1, 2 -> 1, 1 -> 0, 0 -> 3
//     移除 2 -> 1 得到 3 -> 1, 1 -> 0, 0 -> 3  没有入边为0的顶点，无解
//
func canFinish(numCourses int, prerequisites [][]int) bool {
	pres := make([][]int, numCourses)
	incomingEdges := make([]int, numCourses)

	for _, course := range prerequisites {
		id := course[0]
		pre := course[1]
		pres[id] = append(pres[id], pre)
		incomingEdges[pre]++
	}

	queue := []int{}
	for i, count := range incomingEdges {
		if count == 0 {
			queue = append(queue, i)
		}
	}

	res := 0
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		res++

		for _, pre := range pres[id] {
			incomingEdges[pre]--
			if incomingEdges[pre] == 0 {
				queue = append(queue, pre)
			}
		}
	}

	return res == numCourses
}

func test(numCourses int, prerequisites [][]int, expect bool) {
	actual := canFinish(numCourses, prerequisites)
	if actual != expect {
		fmt.Printf("fail! expect %t, got %t\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(2, [][]int{[]int{1, 0}}, true)
	test(2, [][]int{[]int{1, 0}, []int{0, 1}}, false)
	test(3, [][]int{[]int{2, 1}, []int{1, 0}, []int{0, 2}}, false)
	test(3, [][]int{[]int{1, 0}, []int{1, 2}, []int{0, 1}}, false)
	test(3, [][]int{[]int{1, 0}, []int{2, 1}}, true)
}
