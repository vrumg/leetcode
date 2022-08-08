package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	total := 0
	pFast := head

	for {
		pFast = pFast.Next
		total++

		if pFast == nil {
			break
		}
		pFast = pFast.Next
		total++

		if pFast == nil {
			break
		}

		head = head.Next
	}

	if total%2 == 0 {
		return head.Next
	}

	return head
}
