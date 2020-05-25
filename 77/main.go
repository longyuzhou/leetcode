package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	k = min(n, k)
	dfs(n, 1, []int{}, func(nums []int) bool {
		size := len(nums)
		if size == k {
			res = append(res, append([]int{}, nums...))
		}
		return size < k
	})
	return res
}

func dfs(n, start int, path []int, f func(path []int) bool) {
	if !f(path) || start > n {
		return
	}

	for i := start; i <= n; i++ {
		dfs(n, i+1, append(path, i), f)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(combine(0, 0))
	fmt.Println(combine(6, 4))
}
