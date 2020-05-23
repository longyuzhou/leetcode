package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return -1 ^ (-1 << 31)
	}

	x := abs(dividend)
	y := abs(divisor)
	times := 0
	n := 0
	for x >= y {
		if y<<n <= x {
			if n == 0 {
				times++
			} else {
				times += 1<<n - 1<<(n-1)
			}
			n++
		} else {
			x = x - y<<(n-1)
			n = 0
		}
	}
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		times = -times
	}
	return times
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func test(dividend, divisor, expect int) {
	fmt.Println("Input:", dividend, divisor)
	fmt.Println("Expect:", expect)
	actual := divide(dividend, divisor)
	fmt.Println("Actutal:", actual)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
	fmt.Println("====================")
}

func main() {
	test(10, 3, 3)
	test(7, -3, -2)
	test(-2147483648, -1, 2147483647)
}
