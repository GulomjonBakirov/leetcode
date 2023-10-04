package solution

import "fmt"

func winnerOfGame(colors string) bool {
	count := 0

	for i := 1; i < len(colors)-1; i++ {
		prev := colors[i-1]
		curr := colors[i]
		next := colors[i+1]
		fmt.Println(prev, curr, next)
		if prev == 65 && curr == 65 && next == 65 {
			count++
		}
		if prev == 66 && curr == 66 && next == 66 {
			count--
		}

		fmt.Println(count)

	}
	return count > 0
}
