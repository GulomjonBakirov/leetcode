package solution

import "testing"

func TestCountBits(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{
			input: 2,
			want:  []int{0, 1, 1},
		}, {
			input: 5,
			want:  []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, test := range tests {
		got := countBits(test.input)

		for i, n := range got {
			if n != test.want[i] {
				t.Fatalf("Expected %v, but got %v", test.want, got)
			}
		}

	}
}
