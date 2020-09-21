package main

import (
	"strconv"
	"strings"
)

// 67. Add Binary, https://leetcode.com/problems/add-binary/

func addBinary(a string, b string) string {
	carry := 0
	var builder strings.Builder

	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		builder.WriteString(strconv.Itoa(sum % 2))
		carry = sum / 2
	}
	if carry != 0 {
		builder.WriteString(strconv.Itoa(carry))
	}

	return reverse(builder.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
