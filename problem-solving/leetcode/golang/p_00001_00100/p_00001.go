package p_00001_00100

// 1. Two Sum, https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	inverted := make(map[int]int)

	for i, num := range nums {
		idx, exists := inverted[target-num]
		if exists {
			return []int{idx, i}
		} else {
			inverted[num] = i
		}
	}
	return nil
}
