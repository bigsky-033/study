package p_00601_00700

import "sort"

// 645. Set Mismatch, https://leetcode.com/problems/set-mismatch/

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	missed := 1
	duplicated := -1
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == 0 {
			duplicated = nums[i]
		} else if diff > 1 {
			missed = nums[i] + 1
		}
	}
	if nums[len(nums)-1] != len(nums) {
		missed = len(nums)
	}
	return []int{
		duplicated,
		missed,
	}
}
