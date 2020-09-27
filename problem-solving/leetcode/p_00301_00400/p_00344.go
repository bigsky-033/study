package p_00301_00400

// 344. Reverse String, https://leetcode.com/problems/reverse-string/

// Recursive version
func reverseString(s []byte) {
	doReverseString(s, 0, len(s)-1)
}

func doReverseString(s []byte, left int, right int) {
	if left >= right {
		return
	}
	s[left], s[right] = s[right], s[left]
	left++
	right--
	doReverseString(s, left, right)
}
