package p00383ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	mm := buildCharCountMap(magazine)

	for _, char := range ransomNote {
		if v, ok := mm[char]; !ok {
			return false
		} else {
			if v <= 0 {
				return false
			}
			mm[char] = v - 1
		}
	}

	return true
}

func buildCharCountMap(s string) map[rune]int {
	cm := make(map[rune]int, 26) // 26 is number of lower case letters in english
	for _, char := range s {
		if v, ok := cm[char]; ok {
			cm[char] = v + 1
		} else {
			cm[char] = 1
		}
	}
	return cm
}
