package leetcode

func kWeakestRows(mat [][]int, k int) []int {
	weakList := make([]int, 0, k)

	rowCount := len(mat)
	colCount := len(mat[0])

	rowsToIterate := make([]int, rowCount)
	for i := 0; i < rowCount; i++ {
		rowsToIterate[i] = i
	}

	// iterate by columns
	for col := 0; col < colCount; col++ {
		// for each column, get each row by the value at that column
		for i := 0; i < len(rowsToIterate); i++ {
			row := rowsToIterate[i]
			if mat[row][col] == 0 {
				weakList = append(weakList, row)
				if i < len(rowsToIterate) {
					// remove the row from the list, to avoid iterating it again
					// if it is not the last element
					// move index i left by one, to avoid skipping the next element
					rowsToIterate = append(rowsToIterate[:i], rowsToIterate[i+1:]...)
					i--
				}
			}
			if len(weakList) == k {
				return weakList[:k]
			}
		}
	}

	// populate the weakList with the remaining rows by order of row index
	size := len(weakList)
	if size < k {
		for i := 0; i < k-size; i++ {
			weakList = append(weakList, rowsToIterate[i])
		}
	}

	return weakList
}
