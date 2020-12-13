package p_00401_00500

import (
	"sort"
	"strings"
)

// 451. Sort Characters By Frequency, https://leetcode.com/problems/sort-characters-by-frequency/

func frequencySort(s string) string {
	if len(s) < 1 {
		return s
	}
	countsByChar := make(map[rune]int)
	for _, r := range s {
		if count, ok := countsByChar[r]; ok {
			countsByChar[r] = count + 1
		} else {
			countsByChar[r] = 1
		}
	}

	inverted := make(map[int][]rune)
	for r, c := range countsByChar {
		if runes, okay := inverted[c]; okay {
			inverted[c] = append(runes, r)
		} else {
			inverted[c] = []rune{r}
		}
	}

	var counts []int
	for c := range inverted {
		counts = append(counts, c)
	}

	sort.Ints(counts)
	var res strings.Builder
	for i := len(counts) - 1; i >= 0; i-- {
		count := counts[i]
		runes := inverted[count]
		for _, r := range runes {
			for i := 0; i < count; i++ {
				res.WriteRune(r)
			}
		}
	}
	return res.String()
}
