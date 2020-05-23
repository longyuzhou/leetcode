package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}

func test(input []int, expect int) {
	fmt.Println("Input:", input)
	fmt.Println("Expect:", expect)
	actual := removeDuplicates(input)
	fmt.Println("Actutal:", actual)
	fmt.Println("Output:", input)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
	fmt.Println("====================")
}

func main() {
	test([]int{1, 2}, 2)
	test([]int{1, 1, 2}, 2)
	test([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5)
}
