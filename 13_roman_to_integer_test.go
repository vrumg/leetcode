package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		input := "III"
		waited := 3
		resp := romanToInt(input)
		assert.EqualValues(t, waited, resp)
	})

	t.Run("first", func(t *testing.T) {
		input := "LVIII"
		waited := 58
		resp := romanToInt(input)
		assert.EqualValues(t, waited, resp)
	})

	t.Run("first", func(t *testing.T) {
		input := "MCMXCIV"
		waited := 1994
		resp := romanToInt(input)
		assert.EqualValues(t, waited, resp)
	})
}
