package p_00001_00100

import (
	"sort"
	"strings"
)

// 12. Integer to Roman, https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	romanByNumber := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	numbers := make([]int, 0)
	for k := range romanByNumber {
		numbers = append(numbers, k)
	}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	var res strings.Builder
	doIntToRoman(num, numbers, romanByNumber, &res)
	return res.String()
}

func doIntToRoman(num int, numbers []int, romanByNumber map[int]string, result *strings.Builder) {
	if num < 1 {
		return
	}
	for _, n := range numbers {
		q := num / n
		r := num % n

		if q > 0 {
			for i := 0; i < q; i++ {
				result.WriteString(romanByNumber[n])
			}
			doIntToRoman(r, numbers, romanByNumber, result)
			break
		}
	}
}
