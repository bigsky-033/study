package p_01001_01100

import "sort"

// 1051. Height Checker, https://leetcode.com/problems/height-checker/

func heightChecker(heights []int) int {
	copied := make([]int, len(heights))
	copy(copied, heights)
	sort.Ints(copied)

	res := 0
	for i := range heights {
		if heights[i] != copied[i] {
			res++
		}
	}
	return res
}
