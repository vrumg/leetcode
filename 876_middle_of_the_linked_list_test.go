package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	t.Run("ok even", func(t *testing.T) {
		e1 := &ListNode{Val: 1}
		e2 := &ListNode{Val: 2, Next: e1}
		e3 := &ListNode{Val: 3, Next: e2}
		e4 := &ListNode{Val: 5, Next: e3}
		e5 := &ListNode{Val: 4, Next: e4}
		e6 := &ListNode{Val: 3, Next: e5}
		e7 := &ListNode{Val: 2, Next: e6}
		e8 := &ListNode{Val: 1, Next: e7}

		resp := middleNode(e8)
		assert.Equal(t, e4, resp)
	})

	t.Run("ok odd long", func(t *testing.T) {
		e1 := &ListNode{Val: 1}
		e2 := &ListNode{Val: 2, Next: e1}
		e3 := &ListNode{Val: 3, Next: e2}
		e4 := &ListNode{Val: 4, Next: e3}
		e5 := &ListNode{Val: 5, Next: e4}
		e6 := &ListNode{Val: 4, Next: e5}
		e7 := &ListNode{Val: 3, Next: e6}
		e8 := &ListNode{Val: 2, Next: e7}
		e9 := &ListNode{Val: 1, Next: e8}

		resp := middleNode(e9)
		assert.EqualValues(t, e5, resp)
	})

	t.Run("one element", func(t *testing.T) {
		e1 := &ListNode{Val: 1}

		resp := middleNode(e1)
		assert.EqualValues(t, e1, resp)
	})
}
