package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	size := len(s)
	if size < 2 {
		return true
	}

	checked := make([]bool, size)
	for i := 0; i < size; i++ {
		if checked[i] {
			continue
		}
		c1 := s[i]
		c2 := t[i]
		for j := i; j < size; j++ {
			if s[j] == c1 || t[j] == c2 {
				if s[j] != c1 || t[j] != c2 {
					return false
				}
				checked[j] = true
			}
		}
	}
	return true
}

func test(s, t string, expect bool) {
	actutal := isIsomorphic(s, t)
	if expect != actutal {
		fmt.Printf("fail! s: %s, t: %s, expect: %t\n", s, t, expect)
	}
}

func main() {
	test("egg", "add", true)
	test("foo", "bar", false)
	test("paper", "title", true)
}
