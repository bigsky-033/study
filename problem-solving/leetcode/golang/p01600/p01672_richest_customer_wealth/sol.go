package p1672richestcustomerwealth

func maximumWealth(accounts [][]int) int {
	mw := -1

	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		mw = max(mw, sum)
	}

	return mw
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
