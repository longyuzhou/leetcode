package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := [][]int{}
	impl(nums, []int{}, func(p []int) bool {
		if len(p) == len(nums) {
			res = append(res, append([]int{}, p...))
			return false
		}
		return true
	})
	return res
}

func impl(nums []int, p []int, f func(p []int) bool) {
	if !f(p) {
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(p, nums[i]) {
			continue
		}
		impl(nums, append(p, nums[i]), f)
	}
}

func contains(nums []int, x int) bool {
	for _, num := range nums {
		if x == num {
			return true
		}
	}
	return false
}

func do(nums []int) {
	res := permute(nums)
	fmt.Println("Input: ", nums)
	fmt.Println("Output:")
	print2d(res)
}

func main() {
	do([]int{1, 2, 3})
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
