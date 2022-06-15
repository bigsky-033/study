package main

import "fmt"

func FindRoot(roots []int, x int) int {
	if x == roots[x] {
		return x
	}

	roots[x] = FindRoot(roots, roots[x])
	return roots[x]
}

func UnionAndFindCycle(roots []int, ranks []int, x int, y int) (isCycle bool) {
	rootX := FindRoot(roots, x)
	rootY := FindRoot(roots, y)

	if rootX != rootY {
		if ranks[rootX] < ranks[rootY] {
			roots[rootX] = rootY
		} else if ranks[rootY] < ranks[rootX] {
			roots[rootY] = rootX
		} else {
			roots[rootY] = rootX
			ranks[rootX] = ranks[rootX] + 1
		}
	} else {
		isCycle = true
	}
	return
}

func validTree(n int, edges [][]int) bool {
	if n < 2 {
		return true
	}

	roots := make([]int, n)
	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
		ranks[i] = 1
	}
	for _, edge := range edges {
		res := UnionAndFindCycle(roots, ranks, edge[0], edge[1])
		if res {
			return false
		}
	}

	root0 := roots[0]
	for _, i := range roots {
		if root0 != FindRoot(roots, i) {
			return false
		}
	}
	return true
}

func main() {
	// n := 5
	// edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	// edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}

	// n := 4
	// edges := [][]int{{0, 1}, {2, 3}}

	// n := 3
	// edges := [][]int{{1, 0}, {2, 0}}

	n := 8
	edges := [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {1, 2}, {5, 6}, {3, 4}}

	res := validTree(n, edges)
	fmt.Println(res)
}
