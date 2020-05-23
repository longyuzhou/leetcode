package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}
	prev := 0
	for i, x := range nums1 {
		if i != 0 && x == prev {
			continue
		}
		if binarySearch(nums2, x) != -1 {
			res = append(res, x)
		}
		prev = x
	}
	return res
}

func binarySearch(nums []int, x int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := int(uint(i+j) >> 1)
		if nums[mid] < x {
			i = mid + 1
		} else if nums[mid] > x {
			j = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	//fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	//fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(binarySearch([]int{1, 2, 3}, 3))
}
