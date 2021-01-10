package p_00301_00400

// 389. Find the Difference, https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	memo := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		cnt, ok := memo[s[i]]
		if ok {
			memo[s[i]] = cnt + 1
		} else {
			memo[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if cnt, ok := memo[t[i]]; ok {
			if cnt == 1 {
				delete(memo, t[i])
			} else {
				memo[t[i]] = cnt - 1
			}
		} else {
			return t[i]
		}
	}
	return 0
}
