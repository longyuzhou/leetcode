package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	x := []byte(num1)
	y := []byte(num2)
	return string(doMultiply(&x, &y))
}

func doMultiply(num1, num2 *[]byte) []byte {
	reverse(num1)
	reverse(num2)
	x, y := *num1, *num2

	m, n := len(x), len(y)
	ans := make([]byte, m+n)

	for i := 0; i < m; i++ {
		a := x[i] - '0'
		carry := byte(0)
		offset := i
		for j := 0; j < n; j++ {
			b := y[j] - '0'
			c := a*b + carry + ans[offset]
			carry = c / 10
			ans[offset] = c % 10
			offset++
		}
		for carry > byte(0) {
			c := ans[offset] + carry
			carry = c / 10
			ans[offset] = c % 10
			offset++
		}
	}

	pos := 0
	for i := len(ans) - 1; i > 0; i-- {
		if ans[i] != byte(0) {
			pos = i
			break
		}
	}
	ans = ans[:pos+1]
	for i, b := range ans {
		ans[i] = '0' + b
	}
	reverse(&ans)
	return ans
}

func reverse(bytes *[]byte) {
	x := *bytes
	i, j := 0, len(x)-1
	for i < j {
		x[i], x[j] = x[j], x[i]
		i++
		j--
	}
}

func main() {
	test := func(num1 string, num2 string) {
		x, _ := strconv.Atoi(num1)
		y, _ := strconv.Atoi(num2)
		expect := strconv.Itoa(x * y)
		actual := multiply(num1, num2)
		if actual != expect {
			fmt.Printf("fail! expect %s, got %s\n", expect, actual)
		} else {
			fmt.Println("pass")
		}
	}

	// test("0", "0")
	// test("0", "1")
	// test("2", "3")
	// test("123", "4")
	test("123", "456")
}
