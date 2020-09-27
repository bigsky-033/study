package main

// 1299. Replace Elements with Greatest Element on Right Side, https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/

import "math"

func replaceElements(arr []int) []int {
	max := math.MinInt32
	for i := len(arr) - 1; i >= 0; i-- {
		cur := arr[i]
		arr[i] = max
		if cur > max {
			max = cur
		}
	}
	arr[len(arr)-1] = -1
	return arr
}
