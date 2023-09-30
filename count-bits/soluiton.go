package solution

func countBits(n int) []int {
	result := make([]int, n+1)
	i := 0

	for i < n+1 {
		if i == 0 || i == 1 {
			result[i] = i
		}
		result[i] = intToBit(i, 0)
		i++
	}

	return result
}

func intToBit(n, x int) int {
	for n > 0 {
		x += n % 2
		n /= 2
	}

	return x
}
