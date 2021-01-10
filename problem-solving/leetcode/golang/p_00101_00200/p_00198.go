package p_00101_00200

// 198. House Robber, https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
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
