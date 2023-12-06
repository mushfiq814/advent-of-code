package d1p1

import (
	"regexp"
	"strconv"
)

func TotalScore(lines []string) int {
	totalScore := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		totalScore += ExtractCalibrationValue(line)
	}

	return totalScore
}

func ExtractCalibrationValue(input string) int {
	mustCompile := regexp.MustCompile("[0-9]")
	match := mustCompile.FindAllString(input, -1)

	if match == nil || len(match) < 1 {
		panic("no numbers found in input string")
	}

	digit0 := match[0]
	digit1 := match[len(match)-1]

	val, err := strconv.Atoi(digit0 + digit1)

	if err != nil {
		panic(err)
	}

	return val
}
