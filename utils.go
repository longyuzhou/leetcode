package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinkedListToList(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func ListToLinkedList(nums []int) *ListNode {
	var head *ListNode
	if len(nums) > 0 {
		head = &ListNode{Val: nums[0]}
		p := head
		for i := 1; i < len(nums); i++ {
			p.Next = &ListNode{Val: nums[i]}
			p = p.Next
		}
	}
	return head
}

func SameList(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i, num := range nums1 {
		if num != nums2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello world!")
}
