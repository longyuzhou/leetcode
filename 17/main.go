package main

import "fmt"

func letterCombinations(digits string) []string {
	mapping := make(map[byte]string)
	mapping['2'] = "abc"
	mapping['3'] = "def"
	mapping['4'] = "ghi"
	mapping['5'] = "jkl"
	mapping['6'] = "mno"
	mapping['7'] = "pqrs"
	mapping['8'] = "tuv"
	mapping['9'] = "wxyz"

	res := []string{}
	if len(digits) > 0 {
		dfs(digits, &mapping, 0, []byte{}, func(chs []byte) {
			res = append(res, string(chs))
		})
	}
	return res
}

func dfs(digits string, mapping *map[byte]string, start int, chs []byte, f func(chs []byte)) {
	if start >= len(digits) {
		f(chs)
		return
	}

	letters := (*mapping)[digits[start]]
	for i := 0; i < len(letters); i++ {
		dfs(digits, mapping, start+1, append(chs, letters[i]), f)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
