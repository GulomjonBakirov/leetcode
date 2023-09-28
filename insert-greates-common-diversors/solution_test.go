package solution

import (
	"leetcode/utils"
	"testing"
)

func TestInsertGreatestCommonDivisors(t *testing.T) {
	tests := []struct {
		input *utils.ListNode
		want  *utils.ListNode
	}{
		{
			input: utils.GenerateList([]int{18, 6, 10, 3}),
			want:  utils.GenerateList([]int{18, 6, 6, 2, 10, 1, 3}),
		},
	}

	for _, test := range tests {
		got := insertGreatestCommonDivisors(test.input)

		if !utils.EqualList(got, test.want) {
			t.Fatalf("Expected %v, got %v", utils.PrintList(test.want), utils.PrintList(got))
		}
	}
}
