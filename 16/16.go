package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) <= 3 {
		return sum16(nums)
	}

	sort.Ints(nums)
	retval := sum16(nums[0:3])

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			total := nums[i] + nums[j] + nums[k]
			if total > target {
				k--
			} else if total == target {
				return total
			} else {
				j++
			}
			if abs(target-retval) > abs(target-total) {
				retval = total
			}
		}
	}

	return retval
}

func sum16(nums []int) int {
	retval := 0
	for _, num := range nums {
		retval += num
	}
	return retval
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func test(input []int, target, expect int) {
	actual := threeSumClosest(input, target)
	if actual != expect {
		fmt.Println("Input: ", input)
		fmt.Println("Output: ", actual)
		fmt.Println("Expect: ", expect)
		fmt.Println("==========")
	}
}

func main() {
	test([]int{-1, 2, 1, -4}, 1, 2)
	test([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82, 82)
}
