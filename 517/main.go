package main

import "fmt"

func findMinMoves(machines []int) int {
	total := 0
	size := len(machines)
	for _, count := range machines {
		total += count
	}
	if total%size != 0 {
		return -1
	}

	re := total / size
	ans := 0
	co := 0
	for i := 0; i < size; i++ {
		co += machines[i] - re
		ans = max(ans, max(abs(co), machines[i]-re))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test(machines []int, expect int) {
	fmt.Println("=>", machines, expect)
	actual := findMinMoves(machines)
	if actual != expect {
		fmt.Println("<= fail! expect:", expect, ", got", actual)
	} else {
		fmt.Println("<= pass")
	}
}

func main() {
	test([]int{1, 0, 5}, 3)
	test([]int{0, 3, 0}, 2)
	test([]int{0, 2, 0}, -1)
	test([]int{4, 0, 0, 4}, 2)
	test([]int{0, 0, 10, 0, 0, 0, 10, 0, 0, 0}, 8)
}
