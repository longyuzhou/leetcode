package main

import (
	"fmt"
	"sort"
)

type sortBy [][]int

func (a sortBy) Len() int      { return len(a) }
func (a sortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool {
	return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] < a[j][1])
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Sort(sortBy(intervals))
	for len(intervals) > 0 {
		x := intervals[0]
		intervals = intervals[1:]
		for len(intervals) > 0 {
			y := intervals[0]
			if x[1] >= y[0] { // 重叠
				if y[1] > x[1] {
					x[1] = y[1]
				}
				intervals = intervals[1:]
			} else {
				// 因为已排序，所以后面的肯定也不会重叠
				break
			}
		}
		res = append(res, x)
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{}))
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	fmt.Println(merge([][]int{[]int{1, 4}, []int{4, 5}}))
}
