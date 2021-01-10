package p_00801_00900

// 859. Buddy Strings, https://leetcode.com/problems/buddy-strings/

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	counts := make([]int, 26)
	diffCnt := 0

	prevDiffIdx := -1
	for i := range A {
		counts[A[i]-'a']++
		if A[i] != B[i] {
			diffCnt++

			if diffCnt > 2 {
				return false
			}
			if diffCnt == 2 {
				if (A[i] != B[prevDiffIdx]) || (A[prevDiffIdx] != B[i]) {
					return false
				}
			} else if diffCnt == 1 {
				prevDiffIdx = i
			}
		}
	}
	if diffCnt == 2 {
		return true
	} else if diffCnt == 0 {
		for _, c := range counts {
			if c >= 2 {
				return true
			}
		}
	}
	return false
}
