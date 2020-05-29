package main

import "fmt"

func findMaxLength(nums []int) int {
	sum := 0
	sums := make(map[int]int)
	sums[0] = -1
	maxlen := 0
	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}
		if start, ok := sums[sum]; ok {
			maxlen = max(maxlen, i-start)
		} else {
			sums[sum] = i
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}) == 2)
	fmt.Println(findMaxLength([]int{0, 1, 0}) == 2)
}
