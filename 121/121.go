package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	i, buy, maxProfit := 1, prices[0], 0
	for i < len(prices) {
		if prices[i] <= buy {
			buy = prices[i]
		} else {
			profit := prices[i] - buy
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		i++
	}
	return maxProfit
}

func test(prices []int, expect int) {
	actual := maxProfit(prices)
	if actual != expect {
		fmt.Printf("fail input: %d expect: %d actual: %d\n", prices, expect, actual)
	}
}

func main() {
	test([]int{7, 1, 5, 3, 6, 4}, 5)
	test([]int{7, 6, 4, 3, 1}, 0)
	test([]int{2, 1, 2, 0, 1}, 1)
	test([]int{2, 1, 2, 1, 0, 1, 2}, 2)
}
