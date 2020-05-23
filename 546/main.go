package main

import (
	"fmt"
)

func removeBoxes(boxes []int) int {
	return 0
}

func test(boxes []int, expect int) {
	fmt.Println("=>", boxes, expect)
	actual := removeBoxes(boxes)
	if actual != expect {
		fmt.Println("<= fail! actual:", actual)
	} else {
		fmt.Println("<= pass")
	}
}

func main() {
	test([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}, 23)
}
