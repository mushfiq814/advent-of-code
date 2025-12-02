package d1p2

import (
	"strconv"
)

func TotalScore(lines []string) int {
	dial := 50
	count := 0

	for _, line := range lines {
		steps, _ := strconv.Atoi(line[1:])
		initial_dial := dial

		if line[0:1] == "L" {
			steps = -steps
		}
		dial = dial + steps

		increment := Absolute(dial/100 - initial_dial/100)
		count = count + increment
		if HasPassedZero(initial_dial, dial) {
			count++
		}

		dial = MapToRealDial(dial)
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

func Absolute(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func HasPassedZero(start int, end int) bool {
	return end == 0 || (start < 0 && end > 0) || (start > 0 && end < 0)
}

// 5594 - 6905
