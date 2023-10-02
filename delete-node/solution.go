package solution

import "leetcode/utils"

func deleteNode(node *utils.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
