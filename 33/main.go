package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	stack := [][]int{[]int{0, len(nums) - 1}}
	for len(stack) > 0 {
		size := len(stack) - 1
		item := stack[size]
		stack = stack[:size]

		low, high := item[0], item[1]
		if low > high {
			continue
		}

		if nums[low] < nums[high] { // ordered
			pos := binarySearch(nums, low, high, target)
			if pos != -1 {
				return pos
			}
		} else { // not ordered
			mid := int(uint(low+high) >> 1)
			if nums[mid] == target {
				return mid
			}
			stack = append(stack, []int{low, mid - 1}, []int{mid + 1, high})
		}
	}

	return -1
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
	test := func(nums []int, target int, expect int) {
		actual := search(nums, target)
		if actual != expect {
			fmt.Printf("fail! expect %d, got %d\n", expect, actual)
		} else {
			fmt.Println("pass")
		}
	}

	test([]int{4, 5, 6, 7, 0, 1, 2}, 0, 4)
	test([]int{4, 5, 6, 7, 0, 1, 2}, 3, -1)
	test([]int{1}, 0, -1)
}
