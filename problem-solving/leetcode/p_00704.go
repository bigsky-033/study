package main

// 704. Binary Search, https://leetcode.com/problems/binary-search/
// Good to read: https://leetcode.com/problems/binary-search/discuss/423162/Binary-Search-101

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid
		}
	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}
