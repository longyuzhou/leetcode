package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a string containing only digits,
// restore it by returning all possible valid IP address combinations.
// A valid IP address consists of exactly four integers (each integer
// is between 0 and 255) separated by single points.
// Example:
//   Input: "25525511135"
//   Output: ["255.255.11.135", "255.255.111.35"]
// Solution: dfs(注意0的处理)
func restoreIpAddresses(s string) []string {
	res := []string{}

	var dfs func(start int, parts []string)
	dfs = func(start int, parts []string) {
		if len(parts) == 4 && start >= len(s) {
			res = append(res, strings.Join(parts, "."))
		}
		if len(parts) == 4 || start >= len(s) {
			return
		}

		if s[start] == '0' {
			dfs(start+1, append(parts, "0"))
		} else {
			for i := start; i < len(s); i++ {
				num, _ := strconv.Atoi(s[start : i+1])
				if num <= 255 {
					dfs(i+1, append(parts, strconv.Itoa(num)))
				} else {
					break
				}
			}
		}
	}

	dfs(0, []string{})
	return res
}

func main() {
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("255255255255"))
	fmt.Println(restoreIpAddresses("010010"))
}
