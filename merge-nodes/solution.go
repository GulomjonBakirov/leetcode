package solution

import (
	"leetcode/utils"
)

func mergeNodes(head *utils.ListNode) *utils.ListNode {
	sum := 0
	result := &utils.ListNode{}

	res := result

	for head.Next != nil {
		if head.Next.Val == 0 {
			res.Next = &utils.ListNode{Val: sum}
			res = res.Next
			sum = 0
		} else {
			sum += head.Next.Val
		}
		head.Next = head.Next.Next
	}

	return result.Next
}
