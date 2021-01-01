package p_01600_01700

// 1640. Check Array Formation Through Concatenation, https://leetcode.com/problems/check-array-formation-through-concatenation/

type index struct {
	group int
	order int
}

func canFormArray(arr []int, pieces [][]int) bool {
	if arr == nil {
		return true
	}

	indexesByValue := make(map[int]index)
	for i, p := range pieces {
		for j, v := range p {
			indexesByValue[v] = index{group: i, order: j}
		}
	}

	group := -1
	order := 0
	for _, v := range arr {
		if val, ok := indexesByValue[v]; ok {
			if group != val.group {
				group = val.group
				order = 0
			} else {
				order++
			}

			if order != val.order {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
