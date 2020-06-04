package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	memo := make(map[int]bool)

	var solve func(nums []int, key int, total int) bool
	solve = func(nums []int, key int, total int) bool {
		if outcome, ok := memo[key]; ok {
			return outcome
		}

		size := len(nums)
		if total <= nums[size-1] {
			memo[key] = true
			return true
		}

		for i := 0; i < size; i++ {
			num := nums[i]
			if i < size-1 && total-num <= nums[size-1] {
				break
			}
			copyOfNums := append([]int{}, nums[:i]...)
			copyOfNums = append(copyOfNums, nums[i+1:]...)
			if !solve(copyOfNums, key-1<<(num-1), total-num) {
				memo[key] = true
				return true
			}
		}
		memo[key] = false
		return false
	}

	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 == desiredTotal {
		return maxChoosableInteger%2 == 1
	}

	nums := make([]int, maxChoosableInteger)
	for i := range nums {
		nums[i] = i + 1
	}
	return solve(nums, -1^(-1<<maxChoosableInteger), desiredTotal)
}

func main() {
	fmt.Println(canIWin(10, 11) == false)
	fmt.Println(canIWin(4, 6) == true)
	fmt.Println(canIWin(10, 40) == false)
	fmt.Println(canIWin(10, 0) == true)
	fmt.Println(canIWin(5, 50) == false)
}
