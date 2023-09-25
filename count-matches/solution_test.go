package solution

import "testing"

func TestCountMatches(t *testing.T) {
	tests := []struct {
		input_1 [][]string
		input_2 string
		input_3 string
		want    int
	}{
		{
			input_1: [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}},
			input_2: "color",
			input_3: "silver",
			want:    1,
		},
	}

	for _, test := range tests {
		got := countMatches(test.input_1, test.input_2, test.input_3)

		if got != test.want {
			t.Fatalf("Expected %v, got %v", test.want, got)
		}

	}

}
