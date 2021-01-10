package p_00201_00300

// 213. House Robber II, https://leetcode.com/problems/house-robber-ii/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(helper(nums[:len(nums)-1]), helper(nums[1:len(nums)]))
}

func helper(nums []int) int {
	withBefore := 0
	withOutBefore := 0
	for i := 0; i < len(nums); i++ {
		temp := withOutBefore
		withOutBefore = max(withBefore, withOutBefore)
		withBefore = nums[i] + temp
	}
	return max(withBefore, withOutBefore)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
