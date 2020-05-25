package main

import "fmt"

// 问题：一个升序排列的数组，在某个位置发生了旋转，要求找出某元素是否在该数组中存在
// 和33题解法相同
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
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
				return true
			}
		} else { // not ordered
			mid := int(uint(low+high) >> 1)
			if nums[mid] == target {
				return true
			}
			stack = append(stack, []int{low, mid - 1}, []int{mid + 1, high})
		}
	}

	return false
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
	fmt.Println(search([]int{}, 0) == false)
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0) == true)
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3) == false)
}
