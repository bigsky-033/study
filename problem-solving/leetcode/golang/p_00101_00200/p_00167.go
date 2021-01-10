package p_00101_00200

// 167. Two Sum II - Input array is sorted, https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		res := binarySearch(nums, i+1, len(nums)-1, target-num)
		if res != -1 && res != i {
			return []int{i + 1, res + 1}
		}
	}
	return nil
}

func binarySearch(nums []int, lo, hi, target int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}
