package d3p2

import (
	"fmt"
	"regexp"

	d3lib "example.com/mushfiq814/advent-of-code/aoc-2023/day-3"
)

func TotalScore(lines []string) int {
	totalScore := 0
	Test([]string{"4..", ".*.", "..."})
	return totalScore
}

func Test(s []string) {
	L, R, TB := d3lib.CalculateBoundary(s, 0, 2, 1)
  regex := regexp.MustCompile("[0-9]")

  l := regex.FindIndex([]byte(L))
  r := regex.FindIndex([]byte(R))
  tb := regex.FindAllIndex([]byte(TB), -1)

	fmt.Println(l, r, tb)
}
