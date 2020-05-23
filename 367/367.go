package main

import "fmt"

func isPerfectSquare(num int) bool {
	i := 1
	for i*i <= num {
		i++
	}
	return (i-1)*(i-1) == num
}

func main() {
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}
