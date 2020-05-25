package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i <= len(nums); i++ {
		dfs(nums, 0, []int{}, func(path []int) bool {
			if len(path) == i {
				res = append(res, append([]int{}, path...))
			}
			return len(path) < i
		})
	}
	return res
}

func dfs(nums []int, start int, path []int, f func(path []int) bool) {
	if !f(path) || start >= len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		dfs(nums, i+1, append(path, nums[i]), f)
	}
}

func main() {
	fmt.Println(subsets([]int{}))
	fmt.Println(subsets([]int{1, 2, 3}))
}
