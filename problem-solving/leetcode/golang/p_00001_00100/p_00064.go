package p_00001_00100

// 64. Minimum Path Sum, https://leetcode.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	sums := make([][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				sums[i][j] = min(sums[i-1][j], sums[i][j-1]) + grid[i][j]
			} else if i > 0 {
				sums[i][j] = sums[i-1][j] + grid[i][j]
			} else if j > 0 {
				sums[i][j] = sums[i][j-1] + grid[i][j]
			} else {
				sums[i][j] = grid[i][j]
			}
		}
	}
	return sums[m-1][n-1]
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
