package main

// 70. Climbing Stairs, https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
	memo := make(map[int]int)
	return doClimbStairs(n, memo)
}

func doClimbStairs(n int, memo map[int]int) int {
	if value, ok := memo[n]; ok {
		return value
	}
	// Assume, there will be no negative input for n
	if n < 3 {
		return n
	}
	value := doClimbStairs(n-1, memo) + doClimbStairs(n-2, memo)
	memo[n] = value
	return value
}
