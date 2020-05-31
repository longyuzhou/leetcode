package main

import "fmt"

func findMin(nums []int) int {
	size := len(nums)
	first := nums[0]
	last := nums[size-1]
	if first > last {
		for i := 1; i < size; i++ {
			if nums[i-1] > nums[i] {
				return nums[i]
			}
		}
	}
	return first
}

func main() {
	fmt.Println(findMin([]int{0}) == 0)
	fmt.Println(findMin([]int{3, 4, 5, 0, 1, 2}) == 0)
}
