package p00876middleofthelinkedlist

import "testing"

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name   string
		input  *ListNode
		output *ListNode
	}{
		{
			name:   "example 1",
			input:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			output: &ListNode{3, &ListNode{4, &ListNode{5, nil}}},
		},
		{
			name:   "example 2",
			input:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
			output: &ListNode{4, &ListNode{5, &ListNode{6, nil}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mn := middleNode(test.input)

			for test.output != nil && mn != nil {
				if test.output.Val != mn.Val {
					t.Errorf("value is not matched. expected: %v actual: %v", test.output.Val, mn.Val)
				}
				test.output = test.output.Next
				mn = mn.Next
			}
		})
	}
}
