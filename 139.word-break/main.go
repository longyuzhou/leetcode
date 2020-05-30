package main

import (
	"fmt"
)

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// determine if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note:
//   The same word in the dictionary may be reused multiple times in the segmentation.
//   You may assume the dictionary does not contain duplicate words.
//
// Example 1:
//   Input: s = "leetcode", wordDict = ["leet", "code"]
//   Output: true
//   Explanation: Return true because "leetcode" can be segmented as "leet code".
//
// Example 2:
//   Input: s = "applepenapple", wordDict = ["apple", "pen"]
//   Output: true
//   Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//                Note that you are allowed to reuse a dictionary word.
//
// Solution: 动态规划
func wordBreak(s string, wordDict []string) bool {
	dp := make(map[string]bool)

	var fn func(start, end int) bool
	fn = func(start, end int) bool {
		if start >= end {
			return true
		}

		key := fmt.Sprintf("%d,%d", start, end)
		if res, ok := dp[key]; ok {
			return res
		}

		for _, word := range wordDict {
			size := len(word)
			if start+size <= end && s[start:start+size] == word && fn(start+size, end) {
				dp[key] = true
				return true
			}
		}
		dp[key] = false
		return false
	}

	return fn(0, len(s))
}

func main() {
	fmt.Println(wordBreak("firefly", []string{"fire"}) == false)
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}) == true)
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}) == true)
}
