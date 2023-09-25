package solution

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0

	for _, item := range items {
		if ruleKey == "color" {
			if item[1] == ruleValue {
				result += 1
			}
		} else if ruleKey == "type" {
			if item[0] == ruleValue {
				result += 1
			}
		} else {
			if item[2] == ruleValue {
				result += 1
			}
		}
	}

	return result
}
