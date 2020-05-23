package main

import "fmt"

func removeElement(nums []int, val int) int {
	retval := 0
	for _, num := range nums {
		if num != val {
			nums[retval] = num
			retval++
		}
	}
	return retval
}

func test(nums []int, val, expect int) {
	fmt.Println("Input:", nums, val)
	fmt.Println("Expect:", expect)
	actual := removeElement(nums, val)
	fmt.Println("Actutal:", actual)
	fmt.Println("Output:", nums)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
	fmt.Println("====================")
}

func main() {
	test([]int{3, 2, 2, 3}, 3, 2)
	test([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5)
}
