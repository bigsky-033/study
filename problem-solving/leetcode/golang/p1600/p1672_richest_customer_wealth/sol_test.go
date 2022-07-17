package p1672richestcustomerwealth

import "testing"

func Test_maximumWealth(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{
			name:   "example 1",
			input:  [][]int{{1, 2, 3}, {3, 2, 1}},
			output: 6,
		},
		{
			name:   "example 2",
			input:  [][]int{{1, 5}, {7, 3}, {3, 5}},
			output: 10,
		},
		{
			name:   "example 3",
			input:  [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			output: 17,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := maximumWealth(test.input)

			if actual != test.output {
				t.Errorf("value is not matched. expected: %d actual: %d",
					test.output, actual)
			}
		})
	}
}

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
