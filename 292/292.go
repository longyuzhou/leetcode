package main

import "fmt"

func canWinNim(n int) bool {
	results := []bool{false, true, true, true}
	return results[n%len(results)]
}

func test(n int, expect bool) {
	actual := canWinNim(n)
	if actual != expect {
		fmt.Printf("fail! Input: %d, Expect: %t\n", n, expect)
	}
}

func main() {
	test(4, false)
	test(5, true)
	for i := 0; i < 20; i++ {
		fmt.Println(canWinNim(i + 1))
	}
}
