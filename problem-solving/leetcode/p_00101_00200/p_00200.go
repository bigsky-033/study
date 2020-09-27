package p_00101_00200

// 200. Number of Islands, https://leetcode.com/problems/number-of-islands/

type Pair struct {
	i int
	j int
}

const LAND byte = 49
const WATER byte = 48

var EMPTY = struct{}{}

func numIslands(grid [][]byte) int {
	count := 0
	visited := make(map[Pair]struct{})
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == LAND {
				bfs(grid, i, j, visited)
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]byte, i int, j int, visited map[Pair]struct{}) {
	var queue []Pair
	queue = append(queue, Pair{i, j})

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = EMPTY

		x := node.i
		y := node.j
		grid[x][y] = WATER

		if x-1 >= 0 && grid[x-1][y] == LAND {
			queue = append(queue, Pair{x - 1, y})
		}
		if x+1 < len(grid) && grid[x+1][y] == LAND {
			queue = append(queue, Pair{x + 1, y})
		}
		if y-1 >= 0 && grid[x][y-1] == LAND {
			queue = append(queue, Pair{x, y - 1})
		}
		if y+1 < len(grid[0]) && grid[x][y+1] == LAND {
			queue = append(queue, Pair{x, y + 1})
		}
	}
}
