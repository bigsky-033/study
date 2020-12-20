package p_01400_01500

// 1464. Maximum Product of Two Elements in an Array, https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

import "sort"

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-2] - 1) * (nums[len(nums)-1] - 1)
}
