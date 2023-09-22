package solution

import (
	"leetcode/utils"
	"math"
)

func getDecimalValue(head *utils.ListNode) int {
	sum, length := 0, 0
	cur := head

	for cur != nil {
		length += 1
		cur = cur.Next
	}

	i := length - 1

	for head != nil {
		sum += head.Val * int(math.Pow(2, float64(i)))
		i--
		head = head.Next
	}

	return sum
}
