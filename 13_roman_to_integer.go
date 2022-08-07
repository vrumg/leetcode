package leetcode

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	r := []rune(s)
	var total int

	for i := 0; i < len(r); i++ {
		sum, offset := getNextSum(r[i:])
		total += sum
		i += offset
	}

	return total
}

func getNextSum(r []rune) (int, int) {
	var offset int
	var total int

	if len(r) == 1 {
		total = m[r[0]]
		return total, offset
	}

	if r[0] == 'I' && (r[1] == 'V' || r[1] == 'X') {
		total = m[r[1]] - m[r[0]]
		offset = 1
		return total, offset
	}

	if r[0] == 'X' && (r[1] == 'L' || r[1] == 'C') {
		total = m[r[1]] - m[r[0]]
		offset = 1
		return total, offset
	}

	if r[0] == 'C' && (r[1] == 'D' || r[1] == 'M') {
		total = m[r[1]] - m[r[0]]
		offset = 1
		return total, offset
	}

	total = m[r[0]]
	return total, offset
}
