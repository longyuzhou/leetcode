package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	size := len(nums)
	if size == 0 {
		return
	}

	i := 0
	for j := 0; j < size; j++ {
		if nums[j] != 0 {
			if i != j {
				nums[i] = nums[j]
				nums[j] = 0
			}
			i++
		}
	}
}

func test(nums []int) {
	fmt.Print(nums, "=>")
	moveZeroes(nums)
	fmt.Println(nums)
}

func main() {
	test([]int{0, 1, 0, 3, 12})
	test([]int{2, 1})
}
