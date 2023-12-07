package d3p1

import (
	"regexp"
	"strconv"
	"strings"
)

func TotalScore(lines []string) int {
	totalScore := 0
	goodPartNums := GetGoodPartNums(lines)

	for _, num := range goodPartNums {
		totalScore += num
	}
	return totalScore
}

func GetGoodPartNums(s []string) []int {
	result := []int{}
	mustCompile := regexp.MustCompile("[0-9]+")

	for row, line := range s {
		// find start and end of all numbers in this line
		digitIndices := mustCompile.FindAllIndex([]byte(line), -1)
		for _, digitIndex := range digitIndices {
			// find part number inside this capture
			num, _ := strconv.Atoi(line[digitIndex[0]:digitIndex[1]])

			// calculate boundary for part number and check whether it has a special
			// boundary (containing any char other than digits or .)
			boundary := CalculateBoundary(s, digitIndex[0]-1, digitIndex[1], row)
			hasSpecialBoundary := strings.Count(boundary, ".") != len(boundary)

			if hasSpecialBoundary {
				result = append(result, num)
			}
		}
	}
	return result
}

func CalculateBoundary(s []string, left int, right int, center int) string {
	boundary := ""

	if left >= 0 {
		boundary += s[center][left : left+1]
	}
	if right+1 <= len(s[center]) {
		boundary += s[center][right : right+1]
	}

	for _, r := range []int{center - 1, center + 1} {
		adjustedLeft := left
		adjustedRight := right
		if r >= 0 && r < len(s) {
			if left < 0 {
				adjustedLeft = 0
			}
			if right+1 > len(s[r]) {
				adjustedRight = len(s[r]) - 1
			}
			boundary += s[r][adjustedLeft : adjustedRight+1]
		}
	}

	return boundary
}
