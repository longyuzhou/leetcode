package main

import "fmt"

// 先统计字符串中[a-z]出现的次数，然后对比s、t是否一致
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	start := byte('a')
	size := int(byte('z') - start)

	x := make([]int, size)
	for _, r := range s {
		x[byte(r)-start]++
	}
	y := make([]int, size)
	for _, r := range t {
		y[byte(r)-start]++
	}

	for i := 0; i < size; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func test(s, t string, expect bool) {
	actual := isAnagram(s, t)
	if actual != expect {
		fmt.Printf("fail! Input: %s, %s Expect: %t\n", s, t, expect)
	}
}

func main() {
	test("rat", "car", false)
	test("aacc", "ccac", false)
	test("anagram", "nagaram", true)
}
