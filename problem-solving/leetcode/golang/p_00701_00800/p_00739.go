package p_00701_00800

// 739. Daily Temperatures, https://leetcode.com/problems/daily-temperatures/

import "math"

func dailyTemperatures(T []int) []int {
	answer := make([]int, len(T))
	memo := make([]int, 101)

	for i := range memo {
		// Valid memo's value should be range with 1 <= len(T) <= 30000. math.MaxInt32 is just invalid value.
		memo[i] = math.MaxInt32
	}
	for i := len(T) - 1; i >= 0; i-- {
		nextWarm := math.MaxInt32
		for j := T[i] + 1; j <= 100; j++ {
			if memo[j] < nextWarm {
				nextWarm = memo[j]
			}
		}
		if nextWarm < math.MaxInt32 {
			answer[i] = nextWarm - i
		}
		memo[T[i]] = i
	}
	return answer
}
