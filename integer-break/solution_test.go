package solution

import "testing"

func TestIntegerBreak(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 10,
			want:  36,
		},
		{
			input: 2,
			want:  1,
		},
	}

	for _, test := range tests {
		got := integerBreak(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
