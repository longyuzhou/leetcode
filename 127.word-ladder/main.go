package main

import "fmt"

// Given two words (beginWord and endWord), and a dictionary's word list,
// find the length of shortest transformation sequence from beginWord to endWord,
// such that:
//   1. Only one letter can be changed at a time.
//   2. Each transformed word must exist in the word list.
// Note:
//   Return 0 if there is no such transformation sequence.
//   All words have the same length.
//   All words contain only lowercase alphabetic characters.
//   You may assume no duplicates in the word list.
//   You may assume beginWord and endWord are non-empty and are not the same.
// Example 1:
//   Input:
//   beginWord = "hit",
//   endWord = "cog",
//   wordList = ["hot","dot","dog","lot","log","cog"]
//
//   Output: 5
//
//   Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//   return its length 5.
//
// Solution:
//   (1) 将beginWord加入列表s1, s2 = wordList, steps = 1
//   (2) 当s1, s2均不为空时, 执行:
//       (2.1) 遍历s1中元素, a
//               遍历s2中元素, b
//                 如果a与b中相等或只相差1个字母
//                   如果b == endWord，返回steps + 1
//                   将b从s2中移除
//                   如果a != b, 则将b并添加到s1中
//       (2.2) 除步骤(2.1)中添加至s1中元素外，将其它元素自s1中移除
//       (2.3) steps++
//   (4) 重复执行步骤(2)
func transform(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	steps := 1
	s1 := []string{beginWord}
	s2 := wordList
	for len(s1) > 0 && len(s2) > 0 {
		size1 := len(s1)
		for i := 0; i < size1; i++ {
			a := s1[i]
			for j := 0; j < len(s2); j++ {
				b := s2[j]
				if a == b || transform(a, b) {
					if b == endWord {
						return steps + 1
					}
					if a != b {
						s1 = append(s1, b)
					}
					s2 = append(s2[:j], s2[j+1:]...)
					j--
				}
			}
		}
		s1 = s1[size1:]
		steps++
	}
	return 0
}

func main() {
	fmt.Println(ladderLength(
		"hot",
		"dog",
		[]string{"hot", "dog"},
	) == 0)

	fmt.Println(ladderLength(
		"hit",
		"hot",
		[]string{"hot"},
	) == 2)

	fmt.Println(ladderLength(
		"hit",
		"cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"},
	) == 5)

	fmt.Println(ladderLength(
		"hit",
		"cog",
		[]string{"hot", "dot", "dog", "lot", "log"},
	) == 0)
}
