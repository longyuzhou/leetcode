package main

import (
	"fmt"
)

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return addDigits(sum)
}

func test(num, expect int) {
	actual := addDigits(num)
	if actual != expect {
		fmt.Printf("fail! Input: %d, Expect: %d\n", num, expect)
	}
}

func main() {
	test(3, 3)
	test(38, 2)
}
