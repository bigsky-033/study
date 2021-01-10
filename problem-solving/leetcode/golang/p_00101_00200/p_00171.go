package p_00101_00200

// 171. Excel Sheet Column Number, https://leetcode.com/problems/excel-sheet-column-number/

func titleToNumber(s string) int {
	colNum := 0
	for i := 0; i < len(s); i++ {
		colNum = colNum*26 + int(s[i]-('A'-1))
	}
	return colNum
}
