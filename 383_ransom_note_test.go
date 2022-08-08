package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		note := "aaa"
		magazine := "aaab"

		resp := canConstruct(note, magazine)
		assert.True(t, resp)
	})

	t.Run("true", func(t *testing.T) {
		note := "abcdefghijklmnopqrstuvwxyz"
		magazine := "abcdefghijklmnopqrstuvwxyz"

		resp := canConstruct(note, magazine)
		assert.True(t, resp)
	})

	t.Run("false", func(t *testing.T) {
		note := "aa"
		magazine := "ab"

		resp := canConstruct(note, magazine)
		assert.False(t, resp)
	})

	t.Run("false", func(t *testing.T) {
		note := "abcf"
		magazine := "cba"

		resp := canConstruct(note, magazine)
		assert.False(t, resp)
	})
}

func BenchmarkCanConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		note := "abcdefghijklmnopqrstuvwxyz"
		magazine := "abcdefghijklmnopqrstuvwxyz"
		_ = canConstruct(note, magazine)
	}

}
