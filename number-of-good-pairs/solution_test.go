package solution

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 2, 3, 1, 1, 3},
			want:  4,
		},
		{
			input: []int{1, 1, 1, 1},
			want:  6,
		},
	}

	for _, test := range tests {
		got := numIdenticalPairs(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
