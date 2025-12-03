package d3p1

import (
	"slices"
	"strconv"
	"strings"
)

func TotalScore(lines []string) int {
	score := 0

	for _, line := range lines {
		score += MaximumJoltage(line)
	}

	return score
}

func MaximumJoltage(bank string) int {
	batteries_str := strings.Split(bank, "")

	batteries := []int{}
	for _, battery := range batteries_str {
		b, _ := strconv.Atoi(battery)
		batteries = append(batteries, b)
	}

	slices.SortFunc(batteries, func(a,b int) int {
		return b - a
	})

	return batteries[0] * 10 + batteries[1]
}
