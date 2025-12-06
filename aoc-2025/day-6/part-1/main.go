package d6p1

import (
	"regexp"
	"strconv"
)

func TotalScore(lines []string) int {
	score := 0
	leading_spaces := regexp.MustCompile(`^\s+`)
	trailing_spaces := regexp.MustCompile(`\s+$`)
	delimiting_spaces := regexp.MustCompile(`\s+`)

	grid := [][]string{}
	for _, line := range lines {
		// remove leading and trailing spaces
		line = leading_spaces.ReplaceAllString(line, "")
		line = trailing_spaces.ReplaceAllString(line, "")

		// split on delimiting spaces
		grid = append(grid, delimiting_spaces.Split(line, -1))
	}

	for j := 0; j < len(grid[0]); j++ {
		op := grid[len(grid)-1][j]
		var res int
		if op == "+" {
			res = 0
		} else {
			res = 1
		}
		for i := 0; i < len(grid)-1; i++ {
			number, _ := strconv.Atoi(grid[i][j])
			if op == "+" {
				res += number
			} else {
				res *= number
			}
		}
		score += res
	}

	return score
}
