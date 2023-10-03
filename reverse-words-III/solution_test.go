package solution

import "testing"

func TestReverseWord(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "God Ding",
			want:  "doG gniD",
		},
		{
			input: "Let's take LeetCode contest",
			want:  "s'teL ekat edoCteeL tsetnoc",
		},
	}

	for _, test := range tests {
		got := reverseWords(test.input)

		if got != test.want {
			t.Fatalf("Expected %v, but got %v", test.want, got)
		}
	}
}
