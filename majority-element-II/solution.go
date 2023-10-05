package solution

func majorityElement(nums []int) []int {
	number1, number2, count1, count2 := -1, -1, 0, 0
	ans := []int{}

	for _, n := range nums {
		if n == number1 {
			count1++
		} else if n == number2 {
			count2++
		} else if count1 == 0 {
			number1 = n
			count1++
		} else if count2 == 0 {
			number2 = n
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0

	for _, n := range nums {
		if n == number1 {
			count1++
		} else if n == number2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		ans = append(ans, number1)
	}
	if count2 > len(nums)/3 {
		ans = append(ans, number2)
	}
	return ans
}
