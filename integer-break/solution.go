package solution

func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 0
	dp[2] = 1

	if n > 2 {
		dp[3] = 2
	}

	for i := 4; i <= n; i++ {
		grMax := -1

		for j := 2; j <= i/2; j++ {
			n1 := j
			n2 := i - j

			if dp[n1] > n1 {
				n1 = dp[n1]
			}

			if dp[n2] > n2 {
				n2 = dp[n2]
			}

			if n1*n2 > grMax {
				grMax = n1 * n2

			}
		}
		dp[i] = grMax
	}
	return dp[n]
}
