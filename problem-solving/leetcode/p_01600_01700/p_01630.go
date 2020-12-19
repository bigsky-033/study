package p_01600_01700

// 1630. Arithmetic Subarrays, https://leetcode.com/problems/arithmetic-subarrays/

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var res []bool
	for i := 0; i < len(l); i++ {
		subNums := make([]int, r[i]+1-l[i])
		copy(subNums, nums[l[i]:r[i]+1])
		res = append(res, isArithmeticSubArray(subNums))
	}
	return res
}

func isArithmeticSubArray(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	sort.Ints(nums)
	diff := nums[1] - nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if diff != (nums[i+1] - nums[i]) {
			return false
		}
	}
	return true
}
