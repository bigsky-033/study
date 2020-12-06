package p_00701_00800

// 763. Partition Labels, https://leetcode.com/problems/partition-labels/

type interval struct {
	start int
	end int
}

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{0}
	}

	m := make(map[byte]*interval)

	// chars put characters in the order in which it appeared
	var chars []byte
	for i := 0; i < len(S); i++ {
		c := S[i]
		if val, ok := m[c]; ok {
			val.end = i
		} else {
			m[c] = &interval{start: i, end: i}
			chars = append(chars, c)
		}
	}

	var res []int
	cur := m[chars[0]]
	from, to := cur.start, cur.end
	for i := 1; i < len(chars); i++ {
		next := m[chars[i]]
		if next.start > to {
			res = append(res, to-from+1)
			from = next.start
			to = next.end
		} else if next.end > to {
			to = next.end
		}
	}
	res = append(res, to-from+1)
	return res
}
