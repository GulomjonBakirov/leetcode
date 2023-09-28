package solution

func differenceOfSum(nums []int) int {
	sumN := 0
	sumD := 0

	for _, n := range nums {

		if !isDigit(n) {
			sumD += sumDigits(n)
		} else {
			sumD += n
		}

		sumN += n

	}

	if sumD > sumN {
		return sumD - sumN
	} else {
		return sumN - sumD
	}
}
func isDigit(number int) bool {
	digit := rune(number)
	return digit >= 1 && digit <= 9
}

func sumDigits(number int) int {
	sum := 0

	if number < 0 {
		number = -number
	}

	for number > 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}

	return sum

}
