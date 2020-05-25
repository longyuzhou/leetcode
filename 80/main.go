package main

import "fmt"

// Given a sorted array nums,
// remove the duplicates in-place
// such that duplicates appeared at most twice and return the new length.
// Do not allocate extra space for another array,
// you must do this by modifying the input array in-place
// with O(1) extra memory.

type Range struct {
	start int
	stop  int
}

func removeDuplicates(nums []int) int {
	maxRepeat := 2

	// 返回下一个数字的下标范围
	p := 0
	nextRange := func() *Range {
		if p >= len(nums) {
			return nil
		}

		i := p + 1
		for i < len(nums) && nums[p] == nums[i] {
			i++
		}
		res := &Range{p, min(i-1, len(nums)-1)}
		p = i
		return res
	}

	i := 0 // 已处理好的元素个数
	curr := nextRange()
	for curr != nil {
		count := curr.stop - curr.start + 1 // 当前数字的重复次数
		i += min(count, maxRepeat)          // 超出最大重复次数的部分要丢弃
		next := nextRange()
		if next != nil {
			num := nums[next.start]
			for j := 0; j < next.start-i; j++ { // i和next.start的差距代表需要往前移的次数
				nums[i+j] = num
			}
		}
		curr = next
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums) == 5)
	fmt.Println(nums)

	nums = []int{}
	fmt.Println(removeDuplicates(nums) == 0)
	fmt.Println(nums)

	nums = []int{1, 2}
	fmt.Println(removeDuplicates(nums) == 2)
	fmt.Println(nums)
}
