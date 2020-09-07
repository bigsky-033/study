package main

// 34. Find First and Last Position of Element in Sorted Array, https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// Edge cases
// nums: [], target: 0
// nums: [1], target: 1

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}

	idx := binarySearch(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}
	lo, hi := idx, idx

	for ; lo >= 0; lo-- {
		if lo-1 < 0 || nums[lo-1] != target {
			break
		}
	}
	return []int{lo, hi}
}

func binarySearch(nums []int, target int) int {
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
