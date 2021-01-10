package p_00101_00200

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
