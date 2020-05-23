package main

import (
	"fmt"
	"strings"
)

func myAtoi(str string) int {
	return int(do(str))
}

func do(str string) int32 {
	zero := byte('0')
	nine := byte('9')
	plus := byte('+')
	minus := byte('-')

	str = strings.TrimLeft(str, " ")
	if len(str) == 0 {
		return 0
	}

	i := 0
	negative := false
	if ch := str[i]; ch == plus || ch == minus {
		i++
		negative = ch == minus
	}

	IntMax := int32(1<<31 - 1)

	num := int32(0)
	for ; i < len(str); i++ {
		ch := str[i]
		if zero <= ch && ch <= nine {
			digit := int32(ch - zero)
			if num > IntMax/int32(10) || (num == IntMax/int32(10) && digit > int32(7)) {
				if negative {
					return int32(-1 << 31)
				}
				return IntMax
			}
			num = num*int32(10) + digit
		} else {
			break
		}
	}

	if negative {
		num = -num
	}

	return num
}

func check(str string, ans int) {
	output := myAtoi(str)
	if ans != output {
		fmt.Printf("Input: %v\n", str)
		fmt.Printf("Output: %d\n", output)
		fmt.Printf("Expect: %d\n", ans)
		fmt.Println("Fail")
	}
}

func main() {
	check("   -42", -42)
	check("4193 with words", 4193)
	check("words and 987", 0)
	check("-91283472332", -2147483648)
	check("+", 0)
	check("-", 0)
	check("-   ", 0)
	check("+   ", 0)
	check(" +0123", 123)
	check(" -0123", -123)
	check(" -+0123", 0)
	check("  0000000000012345678", 12345678)
}
