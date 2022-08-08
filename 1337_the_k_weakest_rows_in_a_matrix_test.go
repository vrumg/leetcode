package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKWeakestRows(t *testing.T) {
	t.Run("5x5", func(t *testing.T) {
		mat := [][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}
		k := 3
		resp := kWeakestRows(mat, k)
		assert.EqualValues(t, []int{2, 0, 3}, resp)
	})

	t.Run("10x10", func(t *testing.T) {
		mat := [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		k := 3
		resp := kWeakestRows(mat, k)
		assert.EqualValues(t, []int{0, 1, 2}, resp)
	})

	t.Run("10x10", func(t *testing.T) {
		mat := [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		k := 3
		resp := kWeakestRows(mat, k)
		assert.EqualValues(t, []int{1, 2, 3}, resp)
	})

	t.Run("6x3", func(t *testing.T) {
		mat := [][]int{
			{1, 1, 0},
			{1, 0, 0},
			{1, 0, 0},
			{1, 1, 1},
			{1, 1, 0},
			{0, 0, 0},
		}
		k := 4
		resp := kWeakestRows(mat, k)
		assert.EqualValues(t, []int{5, 1, 2, 0}, resp)
	})
}
