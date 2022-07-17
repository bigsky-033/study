package p0412fizzbuzz

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output []string
	}{
		{
			name:   "example 1",
			input:  3,
			output: []string{"1", "2", "Fizz"},
		},
		{
			name:   "example 2",
			input:  5,
			output: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:  "example 3",
			input: 15,
			output: []string{"1", "2", "Fizz", "4",
				"Buzz", "Fizz", "7", "8",
				"Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := fizzBuzz(test.input)

			if len(actual) != len(test.output) {
				t.Fatalf("actual length: %d expected length: %d", len(actual), len(test.output))
			}
			if !reflect.DeepEqual(actual, test.output) {
				t.Errorf("value is not matched. actual: %v expected: %v", actual, test.output)
			}
		})
	}
}
