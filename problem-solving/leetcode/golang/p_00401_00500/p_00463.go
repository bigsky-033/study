package p_00401_00500

// 463. Island Perimeter, https://leetcode.com/problems/island-perimeter/

func islandPerimeter(grid [][]int) int {
	p := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				p += parameter(i, j, grid)
			}
		}
	}
	return p
}

func parameter(i int, j int, grid [][]int) int {
	p := 0
	if i-1 >= 0 && grid[i-1][j] != 1 {
		p++
	}
	if i+1 < len(grid) && grid[i+1][j] != 1 {
		p++
	}
	if j-1 >= 0 && grid[i][j-1] != 1 {
		p++
	}
	if j+1 < len(grid[i]) && grid[i][j+1] != 1 {
		p++
	}
	return p
}
