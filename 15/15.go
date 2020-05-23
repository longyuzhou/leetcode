package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	var prev *int
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x > 0 {
			break
		}
		if prev != nil && *prev == x {
			continue
		}
		sum(nums, 0-x, i+1, func(y, z int) {
			res = append(res, []int{x, y, z})
		})
		prev = &x
	}

	return res
}

func sum(nums []int, n, start int, collector func(x, y int)) {
	i := start
	j := len(nums) - 1
	for i < j {
		for i > start && nums[i] == nums[i-1] && i < j {
			i++
		}
		for j < len(nums)-1 && nums[j] == nums[j+1] && i < j {
			j--
		}
		if i < j {
			x := nums[i]
			y := nums[j]
			sum := x + y
			if sum > n {
				j--
			} else if sum == n {
				collector(x, y)
				i++
				j--
			} else {
				i++
			}
		}
	}
}

func main() {
	fmt.Println("answer:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("answer:", threeSum([]int{-1, 1}))
	fmt.Println("answer:", threeSum([]int{1, 1, -2}))
	fmt.Println("answer:", threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
	fmt.Println("answer:", threeSum([]int{1}))
	fmt.Println("answer:", threeSum([]int{-4, -2, -1}))
	sum([]int{-1, 0, 1, 2}, 1, 0, func(x, y int) {
		fmt.Println("answer:", x, y)
	})
}
