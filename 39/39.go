package main

import "fmt"

// 候选列表已排序并无重复元素
// 每个元素可以用多次，结果不能重复
func combinationSum(candidates []int, target int) [][]int {
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
	for i := start; i < len(candidates); i++ {
		impl(candidates, i, append(nums, candidates[i]), f)
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
	fmt.Println("Input: candidates =", candidates, "target =", target)
	res := combinationSum(candidates, target)
	fmt.Println("A solution set is:")
	print2d(res)
}

func main() {
	do([]int{2, 3, 6, 7}, 7)
	do([]int{2, 3, 5}, 8)
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
