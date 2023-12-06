package d2p1

import (
	"regexp"
	"strconv"
	"strings"
)

func CalculateGamePossibility(game string) (bool, int) {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14

	gameParts := strings.Split(game, ":")
	gameRegex := regexp.MustCompile("Game ([0-9]+)")
	gameIndexMatches := gameRegex.FindStringSubmatch(gameParts[0])
	if gameIndexMatches == nil {
		return false, 0
	}
	gameIndex, _ := strconv.Atoi(gameIndexMatches[1])

	countRegex := regexp.MustCompile("([0-9]+) (blue|red|green)")
	countMatches := countRegex.FindAllStringSubmatch(gameParts[1], -1)
	for _, match := range countMatches {
		count, _ := strconv.Atoi(match[1])
		color := match[2]

		if color == "red" && count > totalRed {
			return false, gameIndex
		} else if color == "green" && count > totalGreen {
			return false, gameIndex
		} else if color == "blue" && count > totalBlue {
			return false, gameIndex
		}
	}

	return true, gameIndex
}

func TotalScore(lines []string) int {
	totalScore := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		possible, game := CalculateGamePossibility(line)
		if possible {
			totalScore += game
		}
	}

	return totalScore
}
