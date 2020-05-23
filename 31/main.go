package main

import "fmt"

func nextPermutation(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	// 从尾部往前，返回升序列表头部+1
	pos := -1
	for i := size - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			pos = i - 1
			break
		}
	}
	if pos != -1 {
		// 从升序列表头部+1的位置“pos”往后，找到大于nums[pos]且最小列表元素的位置
		target := -1
		for i := pos + 1; i < size; i++ {
			if nums[i] > nums[pos] {
				if target == -1 || nums[i] < nums[target] {
					target = i
				}
			}
		}
		// 交换元素位置
		nums[pos], nums[target] = nums[target], nums[pos]
	}

	// 冒泡排序
	i, j := pos+1, len(nums)-1
	for i < j {
		for k := i + 1; k <= j; k++ {
			if nums[k] > nums[i] {
				nums[i], nums[k] = nums[k], nums[i]
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}
}

func test(nums, expect []int) {
	nextPermutation(nums)
	actual := nums
	match := true
	for i, num := range expect {
		if num != actual[i] {
			match = false
			break
		}
	}
	if !match {
		fmt.Println("fail! actual:", actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test([]int{1}, []int{1})
	test([]int{1, 2}, []int{2, 1})
	test([]int{1, 2, 3}, []int{1, 3, 2})
	test([]int{3, 2, 1}, []int{1, 2, 3})
	test([]int{1, 1, 5}, []int{1, 5, 1})
	test([]int{1, 3, 2}, []int{2, 1, 3})
	test([]int{4, 2, 0, 2, 3, 2, 0}, []int{4, 2, 0, 3, 0, 2, 2})
}
