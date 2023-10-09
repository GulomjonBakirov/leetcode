package solution

import "testing"

func TestSearchRange(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			input:  []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			input:  []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			input:  []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			input:  []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, test := range tests {
		got := searchRange(test.input, test.target)

		for i, n := range got {
			if n != test.want[i] {
				t.Fatalf("Expected %v, but got %v", test.want, got)
			}
		}
	}
}
