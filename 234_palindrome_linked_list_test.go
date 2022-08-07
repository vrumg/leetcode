package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	t.Run("false random", func(t *testing.T) {
		head := &ListNode{Val: 1}
		nexPtr := head
		for i := 2; i < 7; i++ {
			next := &ListNode{Val: i}
			nexPtr.Next = next
			nexPtr = next
		}

		resp := isPalindrome(head)
		assert.False(t, resp)
	})

	t.Run("true even elements", func(t *testing.T) {
		e1 := &ListNode{Val: 1}
		e2 := &ListNode{Val: 2, Next: e1}
		e3 := &ListNode{Val: 3, Next: e2}
		e4 := &ListNode{Val: 4, Next: e3}
		e5 := &ListNode{Val: 4, Next: e4}
		e6 := &ListNode{Val: 3, Next: e5}
		e7 := &ListNode{Val: 2, Next: e6}
		e8 := &ListNode{Val: 1, Next: e7}

		resp := isPalindrome(e8)
		assert.True(t, resp)
	})

	t.Run("true odd elements", func(t *testing.T) {
		e1 := &ListNode{Val: 1}
		e2 := &ListNode{Val: 2, Next: e1}
		e3 := &ListNode{Val: 3, Next: e2}
		e4 := &ListNode{Val: 4, Next: e3}
		e5 := &ListNode{Val: 5, Next: e4}
		e6 := &ListNode{Val: 4, Next: e5}
		e7 := &ListNode{Val: 3, Next: e6}
		e8 := &ListNode{Val: 2, Next: e7}
		e9 := &ListNode{Val: 1, Next: e8}

		resp := isPalindrome(e9)
		assert.True(t, resp)
	})
}
