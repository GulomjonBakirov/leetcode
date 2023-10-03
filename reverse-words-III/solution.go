package solution

func reverseWords(s string) string {
	left := 0
	right := 0
	reversed := ""

	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			right = i
			for right > left {
				reversed += string(s[right-1])
				right -= 1
			}
			reversed += string(s[i])
			left = i + 1
		} else if i == len(s)-1 {
			right = i

			for right > left-1 {
				reversed += string(s[right])
				right -= 1
			}
		}

	}

	return reversed
}
