package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, 0, func(x []int) {
		res = append(res, append([]int{}, x...))
	})
	return res
}

func dfs(nums []int, start int, f func(nums []int)) {
	f(nums)

	// 选一个作为开头
	for i := start; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			skip := false
			for k := i; k < j; k++ {
				if nums[k] == nums[j] { // 之前已经选过了
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			dfs(nums, i+1, f)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2, 2}))
}
