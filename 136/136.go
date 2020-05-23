package main

import "fmt"

func singleNumber(nums []int) int {
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}

func test(nums []int, expect int) {
	actual := singleNumber(nums)
	if actual != expect {
		fmt.Printf("fail input: %d expect: %d actual: %d\n", nums, expect, actual)
	}
}

func main() {
	test([]int{2, 2, 1}, 1)
	test([]int{4, 1, 2, 1, 2}, 4)
}
