package p_00401_00500

// 455. Assign Cookies, https://leetcode.com/problems/assign-cookies/

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) < 1 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	i, j, c := 0, 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			c++
			j++
			i++
		} else {
			j++
		}
	}
	return c
}
