package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			max += profit
		}
	}
	return max
}

func test(prices []int, expect int) {
	actual := maxProfit(prices)
	if actual != expect {
		fmt.Printf("fail input: %d expect: %d actual: %d\n", prices, expect, actual)
	}
}

func main() {
	test([]int{7, 1, 5, 3, 6, 4}, 7)
}
