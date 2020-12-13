package p_01001_01100

// 1046. Last Stone Weight, https://leetcode.com/problems/last-stone-weight/

import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		y := stones[n-1]
		x := stones[n-2]
		stones = stones[:n-2]

		if x != y {
			stones = append(stones, y-x)
		}
	}
	if len(stones) > 0 {
		return stones[0]
	} else {
		return 0
	}
}
