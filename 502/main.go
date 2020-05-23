package main

import (
	"fmt"
)

type pair struct {
	profit  int
	capital int
}

type priorityQueue struct {
	queue   []*pair
	compare func(a, b *pair) int
}

func (q *priorityQueue) add(item *pair) {
	q.queue = append(q.queue, item)
	k := len(q.queue) - 1
	for k > 0 {
		parent := (k - 1) >> 1
		if q.compare(item, q.queue[parent]) >= 0 {
			break
		}
		q.queue[k] = q.queue[parent]
		k = parent
	}
	q.queue[k] = item
}

func (q *priorityQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *priorityQueue) peek() *pair {
	if len(q.queue) > 0 {
		return q.queue[0]
	}
	return nil
}

func (q *priorityQueue) poll() *pair {
	size := len(q.queue)
	if size > 0 {
		result := q.queue[0]
		size--
		x := q.queue[size]
		q.queue = q.queue[:size]

		if size > 0 {
			k := 0
			half := size >> 1
			for k < half {
				child := (k << 1) + 1
				c := q.queue[child]
				right := child + 1
				if right < size && q.compare(c, q.queue[right]) > 0 {
					child = right
					c = q.queue[child]
				}
				if q.compare(x, c) <= 0 {
					break
				}
				q.queue[k] = c
				k = child
			}
			q.queue[k] = x
		}

		return result
	}
	return nil
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	profitQueue := priorityQueue{compare: func(a, b *pair) int {
		return b.profit - a.profit
	}}
	capitalQueue := priorityQueue{compare: func(a, b *pair) int {
		return a.capital - b.capital
	}}

	for i := range Profits {
		profitQueue.add(&pair{Profits[i], Capital[i]})
	}

	k = min(k, len(Profits))
	for k > 0 {
		if profitQueue.isEmpty() && capitalQueue.peek().capital > W {
			break
		}
		for !profitQueue.isEmpty() && profitQueue.peek().capital > W {
			capitalQueue.add(profitQueue.poll())
		}
		for !capitalQueue.isEmpty() && capitalQueue.peek().capital <= W {
			profitQueue.add(capitalQueue.poll())
		}
		if profitQueue.isEmpty() {
			continue
		}
		pair := profitQueue.poll()
		W += pair.profit
		k--
	}

	return W
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func test(k int, W int, Profits []int, Capital []int, expect int) {
	actual := findMaximizedCapital(k, W, Profits, Capital)
	if actual != expect {
		fmt.Println("fail! actual:", actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4)
	test(10, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6)
}
