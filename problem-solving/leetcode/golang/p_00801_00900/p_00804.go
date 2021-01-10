package p_00801_00900

import "strings"

// 804. Unique Morse Code Words, https://leetcode.com/problems/unique-morse-code-words/

func uniqueMorseRepresentations(words []string) int {
	moresCodes := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---",
		".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	ur := make(map[string]struct{})
	empty := struct{}{}
	for _, w := range words {
		var b1 strings.Builder
		for _, c := range w {
			b1.WriteString(moresCodes[int(c-'a')])
		}
		r := b1.String()
		if _, okay := ur[r]; !okay {
			ur[r] = empty
		}
	}
	return len(ur)
}
