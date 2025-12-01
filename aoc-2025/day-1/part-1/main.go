package d1p1

import (
	"strconv"
)

func TotalScore(lines []string) int {
	dial := 50
	count := 0

	for _, line := range lines {
		steps, _ := strconv.Atoi(line[1:])

		if line[0:1] == "L" {
			steps = -steps
		}
		dial = dial + steps

		dial = MapToRealDial(dial)

		if dial == 0 {
			count++
		}
	}

	return count
}

func MapToRealDial(dial int) int {
	dial = dial % 100

	if dial < 0 {
		return dial + 100
	}

	if dial >= 100 {
		return dial - 100
	}

	return dial
}
