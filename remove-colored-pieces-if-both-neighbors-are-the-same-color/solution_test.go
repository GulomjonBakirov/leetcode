package solution

import "testing"

func TestWinnerGame(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "ABBBBBBBAAA",
			want:  false,
		},
		// {
		// 	input: "AA",
		// 	want:  false,
		// },
		// {
		// 	input: "ABBBBBBBAAA",
		// 	want:  false,
		// },
	}

	for _, test := range tests {
		got := winnerOfGame(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
