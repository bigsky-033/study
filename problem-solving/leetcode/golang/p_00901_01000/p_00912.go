package p_00901_01000

// 912. Sort an Array, https://leetcode.com/problems/sort-an-array/

func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	ret := make([]int, len(left)+len(right))
	leftIdx, rightIdx, retIdx := 0, 0, 0

	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] < right[rightIdx] {
			ret[retIdx] = left[leftIdx]
			leftIdx++
		} else {
			ret[retIdx] = right[rightIdx]
			rightIdx++
		}
		retIdx++
	}

	for leftIdx < len(left) {
		ret[retIdx] = left[leftIdx]
		leftIdx++
		retIdx++
	}
	for rightIdx < len(right) {
		ret[retIdx] = right[rightIdx]
		rightIdx++
		retIdx++
	}
	return ret
}
