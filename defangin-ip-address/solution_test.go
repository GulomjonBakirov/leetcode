package solution

import "testing"

func TestDefangingIpAddresss(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1.1.1.1",
			want:  "1[.]1[.]1[.]1",
		},
		{
			input: "255.100.50.0",
			want:  "255[.]100[.]50[.]0",
		},
	}

	for _, test := range tests {
		got := defangIPaddr(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
