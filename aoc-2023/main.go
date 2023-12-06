package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/mushfiq814/advent-of-code/aoc-2023/day-1/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2023/day-1/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2023/day-2/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2023/day-2/part-2"
)

func main() {
	day2_part2(false)
}

func readInput(day int, part int, test bool) []string {
	filepath := "./day-" + strconv.Itoa(day) +
		"/part-" + strconv.Itoa(part) +
		"/input_short.txt"

	if !test {
		filepath = "./day-" + strconv.Itoa(day) +
			"/input.txt"
	}

	file, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}

func day1_part1(test bool) {
	lines := readInput(1, 1, test)
	score := d1p1.TotalScore(lines)
	fmt.Println(score)
}

func day1_part2(test bool) {
	lines := readInput(1, 2, test)
	score := d1p2.TotalScore(lines)
	fmt.Println(score)
}

func day2_part1(test bool) {
	lines := readInput(2, 1, test)
	score := d2p1.TotalScore(lines)
	fmt.Println(score)
}

func day2_part2(test bool) {
	lines := readInput(2, 1, test)
	score := d2p2.TotalScore(lines)
	fmt.Println(score)
}
