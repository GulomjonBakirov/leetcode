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
	}

	for _, test := range tests {
		got := pairSum(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, got %v", test.want, got)
		}
	}
}
