package main

import (
	"errors"
	"fmt"
)

type alg struct {
	ring      string
	key       string
	dp        [][]int
	positions map[byte][]int
}

func (this *alg) dfs(ringIdx, keyIdx int) int {
	if keyIdx >= len(this.key) {
		return 0
	}

	if this.dp[ringIdx][keyIdx] != -1 {
		return this.dp[ringIdx][keyIdx]
	}

	cost := -1 ^ (-1 << 31)
	for _, pos := range this.positions[this.key[keyIdx]] {
		step, _ := min(
			abs(ringIdx-pos),
			abs(ringIdx-(pos+len(this.ring))),
			abs(ringIdx+len(this.ring)-pos),
		)
		cost, _ = min(cost, step+this.dfs(pos, keyIdx+1))
	}
	this.dp[ringIdx][keyIdx] = cost
	return cost
}

func findRotateSteps(ring string, key string) int {
	inst := alg{ring: ring, key: key}

	positions := make(map[byte][]int)
	for i, char := range []byte(ring) {
		val, ok := positions[char]
		if !ok {
			val = []int{}
		}
		positions[char] = append(val, i)
	}
	inst.positions = positions

	dp := make([][]int, len(ring))
	for i := 0; i < len(ring); i++ {
		dp[i] = make([]int, len(key))
		for j := 0; j < len(key); j++ {
			dp[i][j] = -1
		}
	}
	inst.dp = dp

	return inst.dfs(0, 0) + len(key)
}

func min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return -1, errors.New("can't work with empty arguments")
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func test(ring, key string, expect int) {
	actual := findRotateSteps(ring, key)
	if actual != expect {
		fmt.Printf("fail! expect %d, get %d\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test("godding", "g", 1)
	test("godding", "gd", 4)
	test("abcde", "ade", 6)
	test("nyngl", "yyynnnnnnlllggg", 19)
	test("edcba", "abcde", 10)
	test("xrrakuulnczywjs", "jrlucwzakzussrlckyjjsuwkuarnaluxnyzcnrxxwruyr", 204)
	test("bicligfijg", "cgijcjlgiggigigijiiciicjilicjflccgilcflijgigbiifiggigiggibbjbijlbcifjlblfggiibjgblgfiiifgbiiciffgbfl", 277)
}
