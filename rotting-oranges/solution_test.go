package solution

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{input: [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}, want: 2},
		{input: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, want: 4},
		{input: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, want: -1},
		{input: [][]int{{0, 2}}, want: 0},
	}

	for _, test := range tests {
		got := orangesRotting(test.input)

		if got != test.want {
			t.Fatalf("Expected %d, got %d", test.want, got)
		}

	}
}
