package main

import "fmt"

// The gray code is a binary numeral system where two successive values differ in only one bit.
// Given a non-negative integer n representing the total number of bits in the code,
// print the sequence of gray code. A gray code sequence must begin with 0.
// Example:
//   Input: 2
//   Output: [0,1,3,2]
//   Explanation:
//   00 - 0
//   01 - 1
//   11 - 3
//   10 - 2
//
// Solution:
//   [0] +1
//   [0, 1] +2
//   [0, 1, 3, 2] +4
//   [0, 1, 3, 2, 6, 7, 5, 4] +8
//   ...
//
func grayCode(n int) []int {
	res := []int{0}
	max := -1 ^ (-1 << n)
	p := 0
	for len(res) < max+1 {
		size := len(res)
		for i := size - 1; i >= 0; i-- {
			num := res[i] + 1<<p
			if num <= max {
				res = append(res, num)
			}
		}
		p++
	}
	return res
}

func test(n int) {
	nums := grayCode(n)
	if len(nums) != -1^(-1<<n)+1 {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(0)
	test(1)
	test(2)
	test(3)
	test(7)
}
