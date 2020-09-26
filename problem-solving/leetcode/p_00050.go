package main

import "math"

// 50. Pow(x, n), https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == math.MinInt32 {
		x = x * x
		n = n / 2
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x*x, n/2)
	}
}
