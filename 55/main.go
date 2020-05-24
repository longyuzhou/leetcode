package main

import "fmt"

// 如果[0, n-1)不存在为0的元素，则肯定能跳到列尾
// 反之，检查是否能越过为0的位置，如果不能越过，则不能跳到队尾
func canJump(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			can := false
			for j := 0; j < i; j++ {
				if j+nums[j] > i {
					can = true
					break
				}
			}
			if !can {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}) == true)
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}) == false)
	fmt.Println(canJump([]int{3, 0, 0, 0}) == true)
}
