package leetcode

// ListNode singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var pSlow *ListNode
	var pFast *ListNode
	var total int

	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	pSlow = head
	pFast = head
	total++

	// count elements
	// get center
	for {
		// pFast with step 2
		if pFast.Next != nil {
			pFast = pFast.Next
			total++
			if pFast.Next != nil {
				pFast = pFast.Next
				total++
			} else {
				break
			}
		} else {
			break
		}

		// pSlow with step 1
		pSlow = pSlow.Next
	}

	// revert right side, get pNext head
	var previous *ListNode
	newHead := pSlow.Next

	for i := 0; i < total/2; i++ {
		if newHead.Next == nil && total/2 == 1 {
			break
		}

		if newHead.Next == nil {
			newHead.Next = previous
			break
		}

		mem := newHead.Next
		newHead.Next = previous
		previous = newHead
		newHead = mem
	}

	oldHead := head
	for i := 0; i < total/2; i++ {
		if oldHead.Val == newHead.Val {
			oldHead = oldHead.Next
			newHead = newHead.Next
			continue
		} else {
			return false
		}
	}

	return true
}
