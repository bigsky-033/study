package main

// 34. Find First and Last Position of Element in Sorted Array, https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// Edge cases
// nums: [], target: 0
// nums: [1], target: 1

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}

	lo := binarySearchForLo(nums, target, 0, len(nums)-1)
	if lo == -1 {
		return []int{-1, -1}
	}
	hi := binarySearchForHi(nums, target, lo, len(nums)-1)
	return []int{lo, hi}
}

func binarySearchForLo(nums []int, target int, lo int, hi int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
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

func binarySearchForHi(nums []int, target int, lo int, hi int) int {
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
