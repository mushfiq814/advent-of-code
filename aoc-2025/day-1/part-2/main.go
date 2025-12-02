package d1p2

import (
	"strconv"
)

func TotalScore(lines []string) int {
	dial := 50
	count := 0

	for _, line := range lines {
		steps, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			steps = -steps
		}

		initial_dial := dial
		dial += steps

		count += Absolute(dial / 100)

		if (initial_dial > 0 && dial <= 0) || (initial_dial == 0 && dial == 0) {
			count++
		}

		dial = MapToRealDial(dial)
	}

	return count
}

func Absolute(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
