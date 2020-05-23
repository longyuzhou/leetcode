package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern, str string) bool {
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		x := pattern[i]
		y := words[i]
		for j := i + 1; j < len(pattern); j++ {
			if (x == pattern[j]) != (y == words[j]) {
				return false
			}
		}
	}

	return true
}

func test(pattern, str string, expect bool) {
	actual := wordPattern(pattern, str)
	if actual != expect {
		fmt.Printf("fail! Input: %s, %s, Expect: %t\n", pattern, str, expect)
	}
}

func main() {
	test("abba", "dog cat cat dog", true)
	test("abba", "dog cat cat fish", false)
	test("abba", "dog dog dog dog", false)
}
