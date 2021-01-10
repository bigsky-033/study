package p_00401_00500

// 415. Add Strings, https://leetcode.com/problems/add-strings/

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	cnt := 0

	var res []int
	n := 0
	for i >= 0 || j >= 0 {
		sum := 0
		sum += n
		if i >= 0 {
			sum += int(num1[i] - '0')
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
		}
		res = append(res, sum%10)
		n = sum / 10

		i--
		j--
		cnt++
	}
	if n > 0 {
		res = append(res, n)
	}
	var b strings.Builder
	for i := len(res) - 1; i >= 0; i-- {
		b.WriteString(strconv.Itoa(res[i]))
	}
	return b.String()
}
