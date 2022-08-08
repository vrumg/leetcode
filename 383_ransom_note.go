package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	// get elements of magazine
	letters := make(map[int32]int)
	for _, letter := range ransomNote {
		letters[letter] += 1
	}

	for _, letter := range magazine {
		if _, ok := letters[letter]; !ok {
			continue
		}
		letters[letter] -= 1
		if letters[letter] == 0 {
			delete(letters, letter)
		}
	}

	if len(letters) != 0 {
		return false
	}

	return true
}
