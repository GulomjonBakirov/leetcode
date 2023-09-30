package solution

import "testing"

func TestPowerOfFour(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 16,
			want:  true,
		},
		{
			input: 0,
			want:  false,
		},
		{
			input: 5,
			want:  false,
		},
	}

	for _, test := range tests {
		got := isPowerOfFour(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
