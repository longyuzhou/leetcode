package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	return nSum(nums, target, 4)
}

func group(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := [][]int{}
	i := 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		res = append(res, nums[i:j])
		i = j
	}
	return res
}

func nSum(nums []int, target int, n int) [][]int {
	groups := group(nums)

	res := [][]int{}
	dfs(groups, 0, n, []int{}, 0, func(total int, path []int) bool {
		if len(path) > 0 && total+path[len(path)-1]*(n-len(path)) > target {
			return false
		}

		if len(path) < n {
			return true
		}

		if total == target {
			tmp := []int{}
			for _, num := range path {
				tmp = append(tmp, num)
			}
			res = append(res, tmp)
		}
		return false
	})
	return res
}

func dfs(groups [][]int, start int, n int, path []int, total int, f func(total int, path []int) bool) {
	if !f(total, path) || start >= len(groups) {
		return
	}

	group := groups[start]
	for i := 0; i <= len(group) && i+len(path) <= n; i++ {
		dfs(groups, start+1, n, append(path, group[:i]...), total+group[0]*i, f)
	}
}

func main() {
	fmt.Println(fourSum([]int{}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, -1))
	fmt.Println(fourSum([]int{-479, -472, -469, -461, -456, -420, -412, -403, -391, -377, -362, -361, -340, -336, -336, -323, -315, -301, -288, -272, -271, -246, -244, -227, -226, -225, -210, -194, -190, -187, -183, -176, -167, -143, -140, -123, -120, -74, -60, -40, -39, -37, -34, -33, -29, -26, -23, -18, -17, -11, -9, 6, 8, 20, 29, 35, 45, 48, 58, 65, 122, 124, 127, 129, 145, 164, 182, 198, 199, 206, 207, 217, 218, 226, 267, 274, 278, 278, 309, 322, 323, 327, 350, 361, 372, 376, 387, 391, 434, 449, 457, 465, 488}, 1979))
}
