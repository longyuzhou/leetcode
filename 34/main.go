package main

import "fmt"

func searchRange(nums []int, target int) []int {
	pos := binarySearch(nums, 0, len(nums)-1, target)
	i, j := pos, pos
	if pos != -1 {
		for i > 0 && nums[i-1] == nums[i] {
			i--
		}
		for j < len(nums)-1 && nums[j] == nums[j+1] {
			j++
		}
	}
	return []int{i, j}
}

func binarySearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := int(uint(low+high) >> 1)
	if nums[mid] < target {
		return binarySearch(nums, mid+1, high, target)
	} else if nums[mid] == target {
		return mid
	}
	return binarySearch(nums, low, mid-1, target)
}

func main() {
	test := func(nums []int, target int, expect []int) {
		actual := searchRange(nums, target)
		match := true
		for i, num := range expect {
			if i >= len(actual) || num != actual[i] {
				match = false
				break
			}
		}
		if !match {
			fmt.Printf("fail! expect %d, got %d\n", expect, actual)
		} else {
			fmt.Println("pass")
		}
	}

	test([]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4})
	test([]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1})
}
