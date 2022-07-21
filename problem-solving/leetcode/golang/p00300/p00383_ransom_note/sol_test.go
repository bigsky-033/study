package p00383ransomnote

import "testing"

func Test_numberOfSteps(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output bool
	}{
		{
			name:   "example 1",
			input:  []string{"a", "b"},
			output: false,
		},
		{
			name:   "example 2",
			input:  []string{"aa", "ab"},
			output: false,
		},
		{
			name:   "example 3",
			input:  []string{"aa", "aab"},
			output: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := canConstruct(test.input[0], test.input[1])

			if test.output != actual {
				t.Errorf("value is not matched. expected: %v actual: %v\n",
					test.output, actual)
			}
		})
	}
}
