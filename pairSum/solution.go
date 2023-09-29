package solution

import (
	"leetcode/utils"
)

func pairSum(head *utils.ListNode) int {
	values := []int{}
	max := 0

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	for i := 0; i < len(values); i++ {
		j := len(values) - 1 - i

		if j < i {
			break
		}

		if max < values[j]+values[i] {
			max = values[j] + values[i]
		}
	}

	return max
}
