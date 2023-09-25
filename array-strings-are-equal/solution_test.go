package solution

import (
	"testing"
)

func TestArrayStringsAreEqual(t *testing.T) {
	tests := []struct {
		input_1 []string
		input_2 []string
		want    bool
	}{
		{
			input_1: []string{"ab", "c"},
			input_2: []string{"a", "bc"},
			want:    true,
		},
		{
			input_1: []string{"a", "cb"},
			input_2: []string{"ab", "c"},
			want:    false,
		},
	}

	for _, test := range tests {
		got := arrayStringsAreEqual(test.input_1, test.input_2)

		if test.want != got {
			t.Fatalf("Expected %v, got %v", test.want, got)
		}

	}
}
