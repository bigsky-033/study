package p_00101_00200

// 153. Find Minimum in Rotated Sorted Array, https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
