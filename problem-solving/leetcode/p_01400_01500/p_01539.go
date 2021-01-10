package p_01400_01500

// 1539. Kth Missing Positive Number, https://leetcode.com/problems/kth-missing-positive-number/

func findKthPositive(arr []int, k int) int {
	i, j := 0, 0
	for {
		if j >= len(arr) || i != arr[j] {
			k--
			if k == 0 {
				return i
			}
		} else {
			j++
		}
		i++
	}
}
