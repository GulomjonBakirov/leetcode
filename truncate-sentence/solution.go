package solution

import (
	"strings"
)

func truncateSentence(s string, k int) string {
	strArr := strings.Split(s, " ")
	res := ""

	for i := 0; i < len(strArr); i++ {
		if i == k-1 {
			res += strArr[i]
			break
		}
		res += strArr[i]
		res += " "

	}

	return res
}
