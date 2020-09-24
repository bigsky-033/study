package main

// 119. Pascal's Triangle II, https://leetcode.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	memo := make([][]int, rowIndex+1)
	for i := range memo {
		memo[i] = make([]int, i+1)
	}
	for j := 0; j < rowIndex+1; j++ {
		memo[rowIndex][j] = getValue(rowIndex, j, memo)
	}
	return memo[rowIndex]
}

func getValue(i int, j int, memo [][]int) int {
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	if j == 0 || i == j {
		memo[i][j] = 1
	} else {
		memo[i][j] = getValue(i-1, j-1, memo) + getValue(i-1, j, memo)
	}
	return memo[i][j]
}
