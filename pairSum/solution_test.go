package solution

import (
	"leetcode/utils"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		input *utils.ListNode
		want  int
	}{
		{
			input: utils.GenerateList([]int{5, 4, 2, 1}),
			want:  6,
		},
		{
			input: utils.GenerateList([]int{4, 2, 2, 3}),
			want:  7,
		},
		{
			input: utils.GenerateList([]int{47, 22, 81, 46, 94, 95, 90, 22, 55, 91, 6, 83, 49, 65, 10, 32, 41, 26, 83, 99, 14, 85, 42, 99, 89, 69, 30, 92, 32, 74, 9, 81, 5, 9}),
			want:  182,
		},
	}

	for _, test := range tests {
		got := pairSum(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, got %v", test.want, got)
		}
	}
}
