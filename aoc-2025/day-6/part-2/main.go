package d6p2

import (
	"math"
	"strconv"
)

func TotalScore(lines []string) int {
	score := 0
	var res int
	var op rune

	for j := 0; j < len(lines[0]); j++ {

		op_str := lines[len(lines)-1][j]
		switch op_str {
		case '+':
			score += res
			res = 0
			op = '+'
		case '*':
			score += res
			res = 1
			op = '*'
		}

		y := 0
		pos := 0
		for i := len(lines) - 2; i >= 0; i-- {
			num_str := string(lines[i][j])
			if num_str == " " {
				continue
			}
			d, _ := strconv.Atoi(num_str)
			y += d * int(math.Pow(10, float64(pos)))
			pos++
		}

		if y != 0 {
			if op == '+' {
				res += y
			} else {
				res *= y
			}
		}
	}

	score += res

	return score
}
