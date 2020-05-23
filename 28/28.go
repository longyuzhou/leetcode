package main

import "fmt"

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	for i := 0; i < l1-l2+1; i++ {
		j := 0
		for j < l2 {
			if haystack[i+j] == needle[j] {
				j++
			} else {
				break
			}
		}
		if j == l2 {
			return i
		}
	}
	return -1
}

func test(haystack, needle string, expect int) {
	fmt.Println("Input:", haystack, needle)
	fmt.Println("Expect:", expect)
	actual := strStr(haystack, needle)
	fmt.Println("Actutal:", actual)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
	fmt.Println("====================")
}

func main() {
	test("", "", 0)
	test("a", "a", 0)
	test("hello", "ll", 2)
	test("aaaaa", "bba", -1)
	test("mississippi", "issi", 1)
}
