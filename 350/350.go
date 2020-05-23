package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := []int{}
	start := 0
	for _, x := range nums1 {
		pos := binarySearch(nums2, start, x)
		if pos != -1 {
			res = append(res, x)
			start = pos + 1
		}
	}
	return res
}

func binarySearch(nums []int, start int, x int) int {
	i, j := start, len(nums)
	for i < j {
		mid := int(uint(i+j) >> 1)
		if nums[mid] < x {
			i = mid + 1
		} else if nums[mid] > x {
			j = mid
		} else {
			k := mid
			for ; k >= start; k-- {
				if nums[k] != x {
					break
				}
			}
			return k + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{2, 1}, []int{1, 2}))
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
