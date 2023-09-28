package solution

import (
	"leetcode/utils"
)

func insertGreatestCommonDivisors(head *utils.ListNode) *utils.ListNode {
	var prev *utils.ListNode
	curr := head

	for curr.Next != nil {
		prev = curr
		curr = curr.Next

		divison := findCommonGreatesDivisions(prev.Val, curr.Val)
		prev.Next = &utils.ListNode{Val: divison}
		prev.Next.Next = curr
	}

	return head
}

func findCommonGreatesDivisions(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	rem := a % b

	return findCommonGreatesDivisions(b, rem)
}
