package solution

import "testing"

func TestDifferenceOfSum(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 15, 6, 3},
			want:  9,
		},
		{
			input: []int{1, 2, 3, 4},
			want:  0,
		},
	}

	for _, test := range tests {
		got := differenceOfSum(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
