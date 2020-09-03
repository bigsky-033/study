package main

// 1346. Check If N and Its Double Exist, https://leetcode.com/problems/check-if-n-and-its-double-exist/

func checkIfExist(arr []int) bool {
	var exists = struct{}{}

	m := make(map[int]struct{})
	for _, v := range arr {
		if _, ok := m[v*2]; ok {
			return true
		}
		if v%2 == 0 {
			if _, ok := m[v/2]; ok {
				return true
			}
		}
		m[v] = exists
	}
	return false
}
