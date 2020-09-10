package main

// 33. Search in Rotated Sorted Array, https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	// Find index of the smallest element
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	// Find out which side target resides. Left or right.
	minIdx := lo
	lo, hi = 0, len(nums)-1
	if target >= nums[minIdx] && target <= nums[hi] {
		lo = minIdx
	} else {
		hi = minIdx
	}

	// Do regular binary search
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
