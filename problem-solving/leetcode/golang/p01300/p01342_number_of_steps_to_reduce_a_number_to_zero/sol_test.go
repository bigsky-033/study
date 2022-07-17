package p01342numberofstepstoreduceanumbertozero

import "testing"

func Test_numberOfSteps(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "example 1",
			input:  14,
			output: 6,
		},
		{
			name:   "example 2",
			input:  8,
			output: 4,
		},
		{
			name:   "example 3",
			input:  123,
			output: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := numberOfSteps(test.input)

			if test.output != actual {
				t.Errorf("value is not matched. expected: %d actual: %d",
					test.output, actual)
			}
		})
	}
}
