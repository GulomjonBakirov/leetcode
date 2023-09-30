package solution

import "testing"

func TestPowerOfThree(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 27,
			want:  true,
		},
		{
			input: 0,
			want:  false,
		},
		{
			input: -1,
			want:  false,
		},
	}

	for _, test := range tests {
		got := isPowerOfThree(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
