package p_00001_00100

// 20. Valid Parentheses, https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	var stack []rune

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if stack == nil || len(stack) < 1 {
				return false
			}
			n := len(stack) - 1
			top := stack[n]
			if (r == ')' && top == '(') ||
				(r == '}' && top == '{') ||
				(r == ']' && top == '[') {
				stack = stack[:n]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
