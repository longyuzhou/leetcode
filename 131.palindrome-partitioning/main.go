package main

import "fmt"

// Given a string s,
// partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.
// Example:
//   Input: "aab"
//   Output:
//   [
//     ["aa","b"],
//     ["a","a","b"]
//   ]
//
// Solution: dfs
func partition(s string) [][]string {
	res := [][]string{}

	isPalindrome := func(i, j int) bool {
		if i > j {
			return false
		}
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index >= len(s) {
			res = append(res, append([]string{}, path...))
		}
		for i := index; i < len(s); i++ {
			if !isPalindrome(index, i) {
				continue
			}
			dfs(i+1, append(path, s[index:i+1]))
		}
	}

	dfs(0, []string{})
	return res
}

func test(s string, expect [][]string) {
	actual := partition(s)
	pass := len(expect) == len(actual)
	if pass {
		for i := range expect {
			for j := range expect[i] {
				if expect[i][j] != actual[i][j] {
					pass = false
					break
				}
			}
		}
	}
	if !pass {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test("", [][]string{[]string{}})
	test("aa", [][]string{[]string{"a", "a"}, []string{"aa"}})
	test("ab", [][]string{[]string{"a", "b"}})
	test("aab", [][]string{[]string{"a", "a", "b"}, []string{"aa", "b"}})
}
