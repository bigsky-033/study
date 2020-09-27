package p_00001_00100

// 11. Container With Most Water: https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	lo, hi := 0, len(height)-1
	currentMax := -1
	for lo < hi {
		currentMax = max(currentMax, getArea(height, lo, hi))
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return currentMax
}

func getArea(height []int, lo int, hi int) int {
	return (hi - lo) * min(height[lo], height[hi])
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
