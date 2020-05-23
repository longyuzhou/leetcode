package main

import "fmt"

func isBadVersion(version int) bool {
	return 2 <= version
}

// 二分查找
func firstBadVersion(n int) int {
	i, j := 1, n
	for i < j {
		mid := int(uint(i+j) >> 1)
		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func main() {
	fmt.Println(firstBadVersion(3))
}
