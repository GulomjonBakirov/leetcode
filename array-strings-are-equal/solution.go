package solution

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var w1 string
	var w2 string

	for _, n := range word1 {
		w1 += n
	}

	for _, n := range word2 {
		w2 += n
	}

	return w1 == w2
}
