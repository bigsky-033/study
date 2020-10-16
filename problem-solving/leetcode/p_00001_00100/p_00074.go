package p_00001_00100

// 74. Search a 2D Matrix, https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) > 0 &&
			matrix[i][0] <= target &&
			matrix[i][len(matrix[i])-1] >= target {
			return binarySearch(matrix[i], target)
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo] == target
}
