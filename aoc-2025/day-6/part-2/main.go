package d6p2

import (
	"math"
	"strconv"
)

func TotalScore(lines []string) int {
	score := 0
	var answer int
	var op rune
	rows := len(lines)

	// iterate through columns
	for col := 0; col < len(lines[0]); col++ {
		op_str := lines[rows-1][col]
		switch op_str {
		case '+':
			answer = 0
			op = '+'
		case '*':
			answer = 1
			op = '*'
		}

		num := 0
		decimal_place := 0
		// iterate through rows bottom up
		for i := rows - 2; i >= 0; i-- {
			num_str := string(lines[i][col])

			// ignore if no number
			if num_str == " " {
				continue
			}

			// build number
			digit, _ := strconv.Atoi(num_str)
			num += digit * int(math.Pow(10, float64(decimal_place)))

			// proceed to next decimal place
			decimal_place++
		}

		// ignore if number is zero (default)
		if num == 0 {
			score += answer
			continue
		}

		if op == '+' {
			answer += num
		} else {
			answer *= num
		}
	}

	// final update of score for the end
	score += answer

	return score
}
