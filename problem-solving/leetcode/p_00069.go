package main

// 69. Sqrt(x), https://leetcode.com/problems/sqrtx/

// find the largest number, which num*num <= x
func mySqrt(x int) int {
	if x < 1 {
		return x
	}

	lo, hi := 1, x
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		// (x/mid < mid) == (x < mid*mid)
		if x/mid < mid {
			// x < mid*mid: mid is too big
			hi = mid - 1
		} else {
			lo = mid
		}
	}
	return lo
}
