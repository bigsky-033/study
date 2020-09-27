package p_00501_00600

// 509. Fibonacci Number, https://leetcode.com/problems/fibonacci-number/

func fib(N int) int {
	memo := make(map[int]int)
	return doFib(N, memo)
}

func doFib(N int, memo map[int]int) int {
	if value, ok := memo[N]; ok {
		return value
	}
	if N == 0 || N == 1 {
		return N
	}
	value := fib(N-1) + fib(N-2)
	memo[N] = value
	return value
}
