package p_00001_00100

// 9. Palindrome Number, https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed, remainder, input := 0, x, x
	for x != 0 {
		remainder = x % 10
		reversed = reversed*10 + remainder
		x /= 10
	}
	return reversed == input
}
