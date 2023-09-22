package solution

import (
	"leetcode/utils"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		input *utils.ListNode
		want  *utils.ListNode
	}{
		{
			input: utils.GenerateList([]int{1, 2, 3, 4, 5, 6}),
			want:  utils.GenerateList([]int{4, 5, 6}),
		},
		{
			input: utils.GenerateList([]int{1, 2, 3, 4, 5}),
			want:  utils.GenerateList([]int{3, 4, 5}),
		},
	}

	for _, test := range tests {
		got := middleNode(test.input)

		if !utils.EqualList(got, test.want) {
			t.Fatalf("Expected %v, got %v", utils.PrintList(test.want), utils.PrintList(got))
		}
	}
}
