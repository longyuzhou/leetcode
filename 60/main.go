package main

import "fmt"

func getPermutation(n int, k int) string {
	nums := []int{}
	x := 1
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
		x = x * i
	}
	if k > x {
		return ""
	}

	bytes := []byte{}
	for i := n; i >= 1; i-- {
		x = x / i
		pos := (k - 1) / x
		bytes = append(bytes, byte(nums[pos])+'0')
		nums = append(nums[0:pos], nums[pos+1:]...)
		k = k - (k-1)/x*x
	}
	return string(bytes)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func test(n int, k int, expect string) {
	actual := getPermutation(n, k)
	if actual != expect {
		fmt.Printf("fail! expect %s, got %s\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(1, 1, "1")
	test(3, 2, "132")
	test(3, 3, "213")
	test(4, 9, "2314")
}
