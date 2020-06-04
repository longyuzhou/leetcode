package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProduct(nums []int) int {
	offset := 0
	res := -1 << 31
	size := len(nums)
	for {
		for offset < size && nums[offset] == 0 {
			res = max(res, 0)
			offset++
		}
		if offset >= size {
			break
		}
		start := offset
		for offset < size && nums[offset] != 0 {
			offset++
		}
		end := offset - 1

		if end-start+1 == 1 {
			res = max(res, nums[start])
		} else if end-start+1 == 2 {
			res = max(res, nums[start])
			res = max(res, nums[start+1])
			res = max(res, nums[start]*nums[start+1])
		} else if end-start+1 >= 3 {
			left := nums[start]
			right := nums[end]
			total := left * right
			for i := 1; i < end-start; i++ {
				if left > 0 {
					left *= nums[start+i]
				}
				if right > 0 {
					right *= nums[end-i]
				}
				total *= nums[start+i]
			}
			res = max(res, total)
			res = max(res, total/left)
			res = max(res, total/right)
		}
	}
	return res
}

func test(nums []int, expect int) {
	actual := maxProduct(nums)
	if actual != expect {
		fmt.Printf("fail! expect %d, got %d\n", expect, actual)
		return
	}
	fmt.Println("pass")
}

func main() {
	test([]int{0}, 0)
	test([]int{-2}, -2)
	test([]int{2, 3, -2, 4}, 6)
	test([]int{-2, 0, -1}, 0)
	test([]int{1, 0, -1, 2, 3, -5, -2}, 60)
	test([]int{-5, 2, 4, 1, -2, 2, -6, 3, -1, -1, -1, -2, -3, 5, 1, -3, -4, 2, -4, 6, -1, 5, -6, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1}, 1492992000)
}
