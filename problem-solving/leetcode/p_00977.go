package main

// 977. Squares of a Sorted Array, https://leetcode.com/problems/squares-of-a-sorted-array/

import "sort"

func sortedSquares(A []int) []int {
	for i, num := range A {
		A[i] = num * num
	}
	sort.Ints(A)
	return A
}
