package main

// 905. Sort Array By Parity, https://leetcode.com/problems/sort-array-by-parity/

func sortArrayByParity(A []int) []int {
	evenIdx := 0

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			if A[evenIdx]%2 != 0 {
				A[i], A[evenIdx] = A[evenIdx], A[i]
			}
			evenIdx++
		}
	}
	return A
}
