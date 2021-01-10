package p_00001_00100

// 3. Longest Substring Without Repeating Characters, https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLen := 0
	charSet := make(map[byte]struct{})
	var empty struct{}

	for left, right := 0, 0; left < len(s) && right < len(s); {
		if _, ok := charSet[s[right]]; !ok {
			charSet[s[right]] = empty
			right++
			maxLen = max(maxLen, right-left)
		} else {
			delete(charSet, s[left])
			left++
		}
	}

	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
