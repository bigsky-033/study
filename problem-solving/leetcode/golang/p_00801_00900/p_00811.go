package p_00801_00900

// 811. Subdomain Visit Count, https://leetcode.com/problems/subdomain-visit-count/

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	counts := make(map[string]int)
	for _, cpdomain := range cpdomains {
		s := strings.Fields(cpdomain)
		cnt, _ := strconv.Atoi(s[0])
		domains := strings.Split(s[1], ".")

		for i := len(domains) - 1; i >= 0; i-- {
			d := strings.Join(domains[i:], ".")
			if val, ok := counts[d]; ok {
				counts[d] = val + cnt
			} else {
				counts[d] = cnt
			}
		}
	}

	var res []string
	for d, c := range counts {
		res = append(res, fmt.Sprintf("%d %s", c, d))
	}
	return res
}
