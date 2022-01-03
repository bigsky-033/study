package sol

// p00026_removeDuplicates solves https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func p00026_removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	k := 1
	nj := 1
	for i := 1; i < len(nums); i++ {
		for j := nj; j < len(nums); j++ {
			if nums[j] != nums[i-1] {
				k++
				nums[i] = nums[j]

				nj = j + 1
				for nj < len(nums) && nums[j] == nums[nj] {
					nj++
				}
				break
			}
		}
	}

	return k
}
