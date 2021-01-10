package p_00201_00300

// 238. Product of Array Except Self, https://leetcode.com/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	cur := 1
	for i := 0; i < len(nums); i++ {
		res[i] = cur
		cur *= nums[i]
	}
	cur = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= cur
		cur *= nums[i]
	}
	return res
}
