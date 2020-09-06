package main

// 704. Binary Search, https://leetcode.com/problems/binary-search/
// Good to read: https://leetcode.com/problems/binary-search/discuss/423162/Binary-Search-101

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
