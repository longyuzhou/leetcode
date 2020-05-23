package main

import (
	"fmt"
)

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	count := 0
	notPrimes := make([]byte, n)
	for i := 2; i < n; i++ {
		if notPrimes[i] == 0 {
			count++
			for j := 2; i*j < n; j++ {
				notPrimes[i*j] = 1
			}
		}
	}
	return count
}

func do(n, expect int) {
	fmt.Println("Input:", n)
	fmt.Println("Expect:", expect)
	actual := countPrimes(n)
	fmt.Println("Actutal:", actual)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
	fmt.Println("====================")
}

func main() {
	do(0, 0)
	do(1, 0)
	do(2, 0)
	do(3, 1)
	do(10, 4)
}
