package d2p2

import (
	"regexp"
	"strconv"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func CalculatePowerOfMinimumSet(game string) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	countRegex := regexp.MustCompile("([0-9]+) (blue|red|green)")
	countMatches := countRegex.FindAllStringSubmatch(game, -1)

	for _, match := range countMatches {
		count, _ := strconv.Atoi(match[1])
		color := match[2]

		if color == "red" && count > maxRed {
			maxRed = count
		} else if color == "green" && count > maxGreen {
			maxGreen = count
		} else if color == "blue" && count > maxBlue {
			maxBlue = count
		}
	}

	return maxRed * maxGreen * maxBlue
}

func TotalScore(lines []string) int {
	totalScore := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		totalScore += CalculatePowerOfMinimumSet(line)
	}

	return totalScore
}
