package main

import "fmt"

// 数字 | 二进制 | 1的个数 | 规律
// 0         0       0
// 1         1       1    1 & 0 == 0
// 2        10       1    2 & 1 == 0
// 3        11       2    3 & 2 == 2
// 4       100       1    4 & 3 == 0
// 5       101       2    5 & 4 == 4
// 6       110       2    6 & 5 == 4
// 7       111       3    7 & 6 == 6
// 8      1000       1    8 & 7 == 0
//
// f(0) = 0
// 如果 n & (n-1) == 0，说明 n == pow(2, ?)，即f(n) = 1
// 如果 n & (n-1) == n - 1，说明n与n-1仅相差1个bit，即f(n) = f(n-1) + 1
// 综上所述 f(n) = f(n&(n-1)) + 1
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func test(num int, expect []int) {
	actual := countBits(num)
	pass := len(expect) == len(actual)
	if pass {
		for i := range expect {
			if expect[i] != actual[i] {
				pass = false
				break
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
	test(0, []int{0})
	test(1, []int{0, 1})
	test(5, []int{0, 1, 1, 2, 1, 2})
}
