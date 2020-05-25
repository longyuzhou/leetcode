package main

import (
	"fmt"
	"sort"
)

func sortColors(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	p := 1
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] > p {
			j--
		}
		for i < j && nums[i] < p {
			i++
		}
		if i < j {
			if nums[i] == p && nums[j] == p {
				// 在(i, j)中找到一个!=p的元素，并交换
				// 如果找不到，则说明已排序
				found := false
				for k := i + 1; k < j; k++ {
					if nums[k] != p {
						found = true
						if nums[k] < p {
							swap(&nums, i, k)
							i++
						} else {
							swap(&nums, j, k)
							j--
						}
						break
					}
				}
				if !found {
					return
				}
			} else {
				swap(&nums, i, j)
				if nums[i] != p {
					i++
				}
				if nums[j] != p {
					j--
				}
			}
		}
	}
}

func swap(nums *[]int, i, j int) {
	s := *nums
	s[i], s[j] = s[j], s[i]
}

func test(nums []int) {
	expect := append([]int{}, nums...)
	sort.Ints(expect)
	sortColors(nums)
	actual := nums
	pass := true
	for i := range expect {
		if expect[i] != actual[i] {
			pass = false
			break
		}
	}
	if !pass {
		fmt.Printf("fail! got %v\n", actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test([]int{})
	test([]int{2, 0, 1})
	test([]int{1, 0, 2})
	test([]int{0, 1, 2})
	test([]int{2, 0, 2, 1, 1, 0})
	test([]int{1, 2, 0})
	test([]int{1, 0, 1})
}
