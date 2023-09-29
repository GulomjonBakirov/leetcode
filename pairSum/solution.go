package solution

import (
	"leetcode/utils"
)

func pairSum(head *utils.ListNode) int {
	cur := head
	pairX, length := 0, 0
	pairs := [][2]int{}
	values := []int{}
	max := 0

	for cur != nil {
		length += 1
		values = append(values, cur.Val)
		cur = cur.Next
	}
	cur = head

	for cur != nil {
		pairY := length - 1 - pairX

		pairs = append(pairs, [2]int{pairX, pairY})

		pairX += 1
		cur = cur.Next
	}

	pairs = pairs[:2]
	cur = head

	for _, pair := range pairs {
		if max < values[pair[0]]+values[pair[1]] {
			max = values[pair[0]] + values[pair[1]]
		}
	}

	return max
}

func findPair(n, i int) int {
	return n - 1 - i
}
