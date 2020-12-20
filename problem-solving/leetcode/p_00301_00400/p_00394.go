package p_00301_00400

// 394. Decode String, https://leetcode.com/problems/decode-string/

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stack []string
	for _, v := range s {
		if v == ']' {
			operand := ""
			for len(stack) > 0 {
				// pop
				n := len(stack) - 1
				top := stack[n]
				stack = stack[:n]

				if top == "[" {
					break
				} else {
					operand = top + operand
				}
			}

			kStr := ""
			for len(stack) > 0 {
				// peek
				n := len(stack) - 1
				top := stack[n]
				if _, err := strconv.Atoi(top); err == nil {
					kStr = top + kStr
					// pop
					stack = stack[:n]
				} else {
					break
				}
			}
			k, _ := strconv.Atoi(kStr)
			v := strings.Repeat(operand, k)
			stack = append(stack, v)
		} else {
			stack = append(stack, string(v))
		}
	}

	var res strings.Builder
	for _, s := range stack {
		res.WriteString(s)
	}
	return res.String()
}
