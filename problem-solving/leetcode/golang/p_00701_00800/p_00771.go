package p_00701_00800

// 771. Jewels and Stones, https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	cnts := make(map[byte]int)
	for i := 0; i < len(J); i++ {
		cnts[J[i]] = 0
	}
	for i := 0; i < len(S); i++ {
		if cnt, okay := cnts[S[i]]; okay {
			cnts[S[i]] = cnt + 1
		}
	}
	res := 0
	for _, v := range cnts {
		res += v
	}
	return res
}
