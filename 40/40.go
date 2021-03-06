package main

import (
	"fmt"
	"sort"
)

// 候选列表未排序，可能由重复元素
// 每个元素只能用一次（两个元素相等的情况是允许的），结果不能由重复
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}

	f := func(nums []int) bool {
		total := sum(nums)
		if total >= target {
			if total == target {
				res = append(res, append([]int{}, nums...))
			}
			return false
		}
		return true
	}

	impl(candidates, 0, []int{}, f)

	return res
}

func impl(candidates []int, start int, nums []int, f func(nums []int) bool) {
	if !f(nums) {
		return
	}
	for i := start; i < len(candidates); {
		num := candidates[i]
		impl(candidates, i+1, append(nums, num), f)
		i++
		for i < len(candidates) {
			if num == candidates[i] {
				i++
			} else {
				break
			}
		}
	}
}

func sum(nums []int) int {
	val := 0
	for _, num := range nums {
		val += num
	}
	return val
}

func do(candidates []int, target int) {
	res := combinationSum2(candidates, target)
	fmt.Println("Input: candidates =", candidates, "target =", target)
	fmt.Println("A solution set is:")
	print2d(res)
}

func main() {
	do([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	do([]int{2, 5, 2, 1, 2}, 5)
}

func print2d(s [][]int) {
	fmt.Println("[")
	for i, row := range s {
		fmt.Print("    [")
		for j, v := range row {
			fmt.Print(v)
			if j < len(row)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]")
		if i < len(s)-1 {
			fmt.Print(",")
		}
		fmt.Print("\n")
	}
	fmt.Println("]")
}
