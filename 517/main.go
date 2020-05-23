package main

import "fmt"

type pair struct {
	idx   int
	count int
}

type priorityQueue struct {
	queue   []*pair
	compare func(a, b *pair) int
}

func (q *priorityQueue) addAll(items ...*pair) {
	for _, item := range items {
		q.add(item)
	}
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

func findMinMoves(machines []int) int {
	total := 0
	size := len(machines)
	for _, count := range machines {
		total += count
	}
	if total%size != 0 {
		return -1
	}

	re := total / size
	ans := 0
	co := 0
	for i := 0; i < size; i++ {
		co += machines[i] - re
		ans = max(ans, max(abs(co), machines[i]-re))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func test(machines []int, expect int) {
	fmt.Println("=>", machines, expect)
	actual := findMinMoves(machines)
	if actual != expect {
		fmt.Println("<= fail! expect:", expect, ", got", actual)
	} else {
		fmt.Println("<= pass")
	}
}

func main() {
	test([]int{1, 0, 5}, 3)
	test([]int{0, 3, 0}, 2)
	test([]int{0, 2, 0}, -1)
	test([]int{4, 0, 0, 4}, 2)
	test([]int{0, 0, 10, 0, 0, 0, 10, 0, 0, 0}, 8)
}
