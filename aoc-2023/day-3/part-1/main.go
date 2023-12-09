package d3p1

import (
	"example.com/mushfiq814/advent-of-code/aoc-2023/day-3"
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
			L, R, TB := d3lib.CalculateBoundary(s, digitIndex[0]-1, digitIndex[1], row)
			boundary := L + R + TB
			hasSpecialBoundary := strings.Count(boundary, ".") != len(boundary)

			if hasSpecialBoundary {
				result = append(result, num)
			}
		}
	}
	return result
}
