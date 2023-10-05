package solution

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{3, 2, 3},
			want:  []int{3},
		},
		{
			input: []int{1},
			want:  []int{1},
		},
	}

	for _, test := range tests {
		got := majorityElement(test.input)

		for i, n := range got {
			if n != test.want[i] {
				t.Fatalf("Expected %v, but got %v", test.want, got)
			}
		}

	}
}
