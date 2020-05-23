package main

import "fmt"

func isUgly(x int) bool {
	if x == 1 {
		return true
	}

	nums := []int{2, 3, 5}

	for x > 0 {
		y := divide(x, nums)
		if y == x {
			break
		}
		x = y
	}

	if x != 1 && !contains(nums, x) {
		return false
	}

	return true
}

func divide(x int, nums []int) int {
	for _, num := range nums {
		if x%num == 0 {
			return x / num
		}
	}
	return x
}

func contains(nums []int, x int) bool {
	for _, num := range nums {
		if num == x {
			return true
		}
	}
	return false
}

func test(x int, expect bool) {
	actual := isUgly(x)
	if actual != expect {
		fmt.Printf("fail! Input: %d, Expect: %t\n", x, expect)
	}
}

func main() {
	test(1, true)
	test(6, true)
	test(8, true)
	test(14, false)
}
