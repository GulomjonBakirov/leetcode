package solution

func myPow(x float64, n int) float64 {
	var y float64
	if n == 0 {
		return 1
	}
	y = myPow(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		if n > 0 {
			return x * y * y
		} else {
			return (y * y) / x
		}
	}
}
