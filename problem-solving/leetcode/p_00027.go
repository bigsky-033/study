package main

// 27. Remove Element: https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	pos := 0
	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
