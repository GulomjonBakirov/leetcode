package solution

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	start, end, count := -1, -1, 0

	for i, num := range nums {

		if num == target {
			if count == 0 {
				start = i
				end = i
				count += 1
			}
			end = i
		}

	}

	return []int{start, end}
}
