package d1p2

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
	mustCompile := regexp.MustCompile("([0-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine))")
	match := mustCompile.FindAllString(input, -1)

	if match == nil || len(match) < 1 {
		panic("no numbers found in input string")
	}

	digits := []string{match[0], match[len(match)-1]}
	result := []string{"0", "0"}

	for i, num := range digits {
		switch num {
		case "one":
			result[i] = "1"
		case "two":
			result[i] = "2"
		case "three":
			result[i] = "3"
		case "four":
			result[i] = "4"
		case "five":
			result[i] = "5"
		case "six":
			result[i] = "6"
		case "seven":
			result[i] = "7"
		case "eight":
			result[i] = "8"
		case "nine":
			result[i] = "9"
		default:
			result[i] = num
		}
	}

	val, err := strconv.Atoi(result[0] + result[1])

	if err != nil {
		panic(err)
	}

	return val
}
