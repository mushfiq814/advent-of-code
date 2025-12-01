package d3p2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	d3lib "example.com/mushfiq814/advent-of-code/aoc-2023/day-3"
)

func TotalScore(lines []string) int {
	totalScore := 0
	FindPartNumsInBoundary([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	})
	return totalScore
}

func FindPartNumsInBoundary(s []string) {
	result := []int{}
	mustCompile := regexp.MustCompile("[0-9]+")

	for row, line := range s {
		// find start and end of all numbers in this line
		digitIndices := mustCompile.FindAllIndex([]byte(line), -1)
		for _, digitIndex := range digitIndices {
			// find part number inside this capture
			num, _ := strconv.Atoi(line[digitIndex[0]:digitIndex[1]])
			fmt.Println(num)

			// calculate boundary for part number and check whether it has a special
			// boundary (containing *)
			L, R, TB := d3lib.CalculateBoundary(s, digitIndex[0]-1, digitIndex[1], row)
			boundary := L + R + TB
			hasSpecialBoundary := strings.Count(boundary, "*") > 0

			if hasSpecialBoundary {
				result = append(result, num)
			}
		}
	}

	fmt.Println(result)

	// regex := regexp.MustCompile("[0-9]")
	// for _, index := range regex.FindIndex([]byte(L)) {
	// 	fmt.Println(index)
	// }
	// for _, index := range regex.FindIndex([]byte(R)) {
	// 	fmt.Println(index)
	// }
	// for _, match := range regex.FindAllIndex([]byte(TB), -1) {
	// 	fmt.Println(match)
	// }
}
