package solution

import "testing"

func TestPow(t *testing.T) {
	tests := []struct {
		input_1 float64
		input_2 int
		want    float64
	}{
		{
			input_1: 2.00000,
			input_2: 10,
			want:    1024.00000,
		},
	}

	for _, test := range tests {
		got := myPow(test.input_1, test.input_2)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
