package main

import (
	"fmt"
	"sort"
)

// Given a collection of integers that might contain duplicates, nums,
// return all possible subsets (the power set).
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//   Input: [1,2,2]
//   Output:
//   [
//     [2],
//     [1],
//     [1,2,2],
//     [2,2],
//     [1,2],
//     []
//   ]
//
// Solution: 基于dfs的组合算法，排序后，跳过重复的元素即可
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i <= len(nums); i++ {
		for _, row := range subset(nums, i) {
			res = append(res, row)
		}
	}
	return res
}

func subset(nums []int, n int) [][]int {
	res := [][]int{}

	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		if start >= len(nums) || len(path) == n {
			if len(path) == n {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for i := start; i < len(nums); i++ {
			skip := false
			for j := start; j < i; j++ {
				if nums[i] == nums[j] {
					skip = true
					break
				}
			}
			if !skip {
				dfs(i+1, append(path, nums[i]))
			}
		}
	}

	dfs(0, []int{})
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
