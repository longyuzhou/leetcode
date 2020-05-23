package main

import "fmt"

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

func test(nums []int, target int, expect []int) {
	actual := twoSum(nums, target)
	if actual[0] != expect[0] || actual[1] != expect[1] {
		fmt.Printf("fail input: %d, %d expect: %d actual: %d\n", nums, target, expect, actual)
	}
}

func main() {
	test([]int{2, 7, 11, 15}, 9, []int{1, 2})
}
