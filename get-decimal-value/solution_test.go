package solution

import (
	"leetcode/utils"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		input *utils.ListNode

		want int
	}{
		{
			input: utils.GenerateList([]int{1, 0, 1}),
			want:  5,
		},
		{
			input: utils.GenerateList([]int{0}),
			want:  0,
		},
		{
			input: utils.GenerateList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}),
			want:  18880,
		},
	}

	for _, test := range tests {
		got := getDecimalValue(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, got %v", test.want, got)
		}
	}
}
