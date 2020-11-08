package p_00001_00100

// 49, Group Anagrams, https://leetcode.com/problems/group-anagrams/

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		k := string(runes)
		if a, ok := m[k]; ok {
			m[k] = append(a, s)
		} else {
			m[k] = []string{s}
		}
	}

	ag := make([][]string, 0, len(m))
	for _, v := range m {
		ag = append(ag, v)
	}
	return ag
}
