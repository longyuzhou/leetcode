package main

import "fmt"

type Node struct {
	price int
	count int
	next  *Node
}

type StockSpanner struct {
	head *Node
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	p := this.head
	count := 1
	for p != nil && p.price <= price {
		count += p.count
		p = p.next
	}
	this.head = &Node{price, count, p}
	return count
}

func main() {
	ss := Constructor()
	input := []int{100, 80, 60, 70, 60, 75, 85}
	output := []int{1, 1, 1, 2, 1, 4, 6}
	for i, price := range input {
		actual := ss.Next(price)
		expect := output[i]
		fmt.Printf("input: %d, expect: %d, actual: %d", price, expect, actual)
		if actual != expect {
			fmt.Println(" fail")
			break
		} else {
			fmt.Println(" pass")
		}
	}
}
