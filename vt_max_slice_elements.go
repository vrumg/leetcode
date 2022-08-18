package leetcode

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

func findKMaxElements(k int, slice []int) []int {
	temp := make([]int, 0, k)
	temp = append(temp, slice[:k]...)

	idxMin, min := findMin(temp)

	for idx := k; idx < len(slice); idx++ {
		if slice[idx] > min {
			temp[idxMin] = slice[idx]
			idxMin, min = findMin(temp)
		}
	}

	return temp
}

func findMin(slice []int) (int, int) {
	min := slice[0]
	idxMin := 0

	for idx, _ := range slice {
		if slice[idx] < min {
			min = slice[idx]
			idxMin = idx
		}
	}

	return idxMin, min
}
