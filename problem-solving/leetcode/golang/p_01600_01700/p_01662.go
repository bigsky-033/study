package p_01600_01700

// 1662. Check If Two String Arrays are Equivalent, https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

func arrayStringsAreEqual(words1 []string, words2 []string) bool {
	word1Idx := 0
	char1Idx := 0
	word2Idx := 0
	char2Idx := 0

	for word1Idx < len(words1) && word2Idx < len(words2) {
		if words1[word1Idx][char1Idx] != words2[word2Idx][char2Idx] {
			return false
		}

		if char1Idx < len(words1[word1Idx])-1 {
			char1Idx++
		} else {
			word1Idx++
			char1Idx = 0
		}
		if char2Idx < len(words2[word2Idx])-1 {
			char2Idx++
		} else {
			word2Idx++
			char2Idx = 0
		}
	}
	return word1Idx == len(words1) && word2Idx == len(words2)
}
