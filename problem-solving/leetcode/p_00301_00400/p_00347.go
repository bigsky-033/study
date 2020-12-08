package p_00301_00400

// 347. Top K Frequent Elements, https://leetcode.com/problems/top-k-frequent-elements/

import "sort"

func topKFrequent(nums []int, k int) []int {
	countByNumber := make(map[int]int)
	for _, num := range nums {
		if val, ok := countByNumber[num]; ok {
			countByNumber[num] = val + 1
		} else {
			countByNumber[num] = 1
		}
	}
	inverted := make(map[int][]int)
	for n, c := range countByNumber {
		if val, ok := inverted[c]; ok {
			inverted[c] = append(val, n)
		} else {
			inverted[c] = []int{n}
		}
	}
	var counts []int
	for cnt := range inverted {
		counts = append(counts, cnt)
	}
	sort.Ints(counts)

	var res []int
	for i := 0; i < k && len(res) < k; i++ {
		idx := counts[len(counts)-1-i]
		res = append(res, inverted[idx]...)
	}
	return res
}
