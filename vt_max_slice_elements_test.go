package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/*
задача: Максимальные элементы в массиве
условие: Дан массив чисел nums и некоторое неизвестное науке число k.
Нужно написать функцию, которая вынимает k самых больших чисел из массива nums.
пример:
// ввод
nums := []int{100, 50, 0, 150, 100, 0, -30, 70}
k := 3
// ожидаемый вывод (в любом порядке)
// 100 150 100
*/

func TestFindKMaxElements(t *testing.T) {
	slice := []int{100, 50, 0, 150, 100, 0, -30, 70}

	t.Run("k=3", func(t *testing.T) {
		k := 3
		expected := []int{100, 100, 150}

		res := findKMaxElements(k, slice)
		sort.Slice(res, func(i, j int) bool {
			if res[i] > res[j] {
				return false
			}
			return true
		})
		assert.EqualValues(t, expected, res)
	})

	t.Run("k=1", func(t *testing.T) {
		k := 1
		expected := []int{150}

		res := findKMaxElements(k, slice)
		assert.EqualValues(t, expected, res)
	})

	t.Run("k=5", func(t *testing.T) {
		k := 5
		expected := []int{50, 70, 100, 100, 150}

		res := findKMaxElements(k, slice)
		sort.Slice(res, func(i, j int) bool {
			if res[i] > res[j] {
				return false
			}
			return true
		})
		assert.EqualValues(t, expected, res)
	})
}
