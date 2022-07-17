package p0412fizzbuzz

import (
	"reflect"
	"strconv"
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

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)

	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}

	return res
}
