package p_00001_00100

// 17. Letter Combinations of a Phone Number, https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}
	lettersByDigit := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	solve(digits, 0, "", &result, lettersByDigit)
	return result
}

func solve(digits string, currentPos int, combination string, result *[]string, lettersByDigit map[byte]string) {
	if len(digits) == currentPos {
		*result = append(*result, combination)
	} else {
		digit := digits[currentPos]
		letters := lettersByDigit[digit]
		for i := 0; i < len(letters); i++ {
			solve(digits, currentPos+1, combination+string(letters[i]), result, lettersByDigit)
		}
	}
}
