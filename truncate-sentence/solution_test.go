package solution

import "testing"

func TestTruncateSentence(t *testing.T) {
	tests := []struct {
		input_1 string
		input_2 int
		want    string
	}{
		{
			input_1: "Hello how are you Contestant",
			input_2: 4,
			want:    "Hello how are you",
		},
		{
			input_1: "What is the solution to this problem",
			input_2: 4,
			want:    "What is the solution",
		},
	}

	for _, test := range tests {
		got := truncateSentence(test.input_1, test.input_2)

		if got != test.want {
			t.Fatalf("Expected %v, got %v", test.want, got)
		}
	}
}
