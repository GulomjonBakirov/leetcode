package solution

import (
	"fmt"
	"leetcode/utils"
)

func middleNode(head *utils.ListNode) *utils.ListNode {
	length, i := 0, 0
	cur := head
	res := &utils.ListNode{}

	for cur != nil {
		length++
		cur = cur.Next
	}

	length = length / 2
	result := res

	for head != nil {
		if i >= length {
			result.Next = &utils.ListNode{Val: head.Val}
			fmt.Println(result.Val, head.Val, i, length)
			result = result.Next
		}
		head = head.Next
		i++
	}

	return res.Next
}
