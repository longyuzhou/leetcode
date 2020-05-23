package main

import (
	"fmt"
	"sort"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	n := sort.Search(x, func(n int) bool {
		return n*n > x
	})
	return n - 1
}

func test(x, expect int) {
	actual := mySqrt(x)
	if actual != expect {
		fmt.Printf("fail x=%d expect=%d actual=%d\n", x, expect, actual)
	}
}

func main() {
	test(1, 1)
	test(4, 2)
	test(8, 2)
}
