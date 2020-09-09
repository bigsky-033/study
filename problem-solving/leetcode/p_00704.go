package main

// 704. Binary Search, https://leetcode.com/problems/binary-search/
// Good to read: https://leetcode.com/problems/binary-search/discuss/423162/Binary-Search-101

// Note
// If number of elements are even, there are two choices for mid
// 1) Right/upper mid: lo + (hi-lo+1)/2
// 2) Left/lower mid: lo + (hi-lo)/2

// Two avoid infinity loop,
// 1) Right/upper mid: we have to update like,
// 		if target < nums[mid] {
//			hi = mid - 1 // exclude mid
//		} else {
//			lo = mid
//		}

// 2) Left/lower mid: we have to update like,
//		if target > nums[mid] {
//			lo = mid + 1 // exclude mid
//		} else {
//			hi = mid
//		}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo+1)/2 // right/upper mid
		if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid
		}

		if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}

	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}
