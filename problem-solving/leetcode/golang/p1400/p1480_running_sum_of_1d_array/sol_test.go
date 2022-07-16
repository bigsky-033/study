package p1480runningsumof1darray

import "testing"

func Test_runningSum(t *testing.T) {
	tests := []struct {
		name string
		input []int
		output []int
	} {
		{
			name: "example 1",
			input: []int{1,2,3,4},
			output: []int{1,3,6,10},
		},
		{
			name: "example 2",
			input: []int{1,1,1,1,1},
			output: []int{1,2,3,4,5},
		},
		{
			name: "example 3",
			input: []int{3,1,2,10,1},
			output: []int{3,4,6,16,17},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T)  {
			actual := runningSum(test.input)

			if len(actual) != len(test.output) {
				t.Fatalf("actual length: %d expected length: %d", len(actual), len(test.output))
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != test.output[i] {
					t.Fatalf("value is not matched. actual[%d]: %d, expected[%d]: %d\n", i, actual[i], i, test.output[i])
				}
			}
		})
	}
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
