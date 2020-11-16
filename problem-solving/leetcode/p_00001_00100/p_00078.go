package p_00001_00100

// 78. Subsets, https://leetcode.com/problems/subsets/

func subsets(nums []int) [][]int {
	var ss [][]int
	var curr []int
	backtrack(nums, 0, &curr, &ss)
	return ss
}

func backtrack(nums []int, idx int, current *[]int, result *[][]int) {
	copied := make([]int, len(*current))
	copy(copied, *current)
	*result = append(*result, copied)

	for i := idx; i < len(nums); i++ {
		*current = append(*current, nums[i])
		backtrack(nums, i+1, current, result)
		*current = (*current)[:len(*current)-1]
	}
}
