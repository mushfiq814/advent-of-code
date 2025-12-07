package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/mushfiq814/advent-of-code/aoc-2025/day-1/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-1/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-2/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-2/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-3/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-3/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-4/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-4/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-5/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-5/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-6/part-1"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-6/part-2"
	"example.com/mushfiq814/advent-of-code/aoc-2025/day-7/part-1"
)

func main() {
	day1_part1(false)
	day1_part2(false)
	day2_part1(false)
	day2_part2(false)
	day3_part1(false)
	day3_part2(false)
	day4_part1(false)
	day4_part2(false)
	day5_part1(false)
	day5_part2(false)
	day6_part1(false)
	day6_part2(false)
	day7_part1(false)
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

	// remove ending newline
	lines := strings.Split(string(file), "\n")
	return lines[:len(lines)-1]
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
	lines := readInput(2, 2, test)
	score := d2p2.TotalScore(lines)
	fmt.Println(score)
}

func day3_part1(test bool) {
	lines := readInput(3, 1, test)
	score := d3p1.TotalScore(lines)
	fmt.Println(score)
}

func day3_part2(test bool) {
	lines := readInput(3, 2, test)
	score := d3p2.TotalScore(lines)
	fmt.Println(score)
}

func day4_part1(test bool) {
	lines := readInput(4, 1, test)
	score := d4p1.TotalScore(lines)
	fmt.Println(score)
}

func day4_part2(test bool) {
	lines := readInput(4, 2, test)
	score := d4p2.TotalScore(lines)
	fmt.Println(score)
}

func day5_part1(test bool) {
	lines := readInput(5, 1, test)
	score := d5p1.TotalScore(lines)
	fmt.Println(score)
}

func day5_part2(test bool) {
	lines := readInput(5, 2, test)
	score := d5p2.TotalScore(lines)
	fmt.Println(score)
}

func day6_part1(test bool) {
	lines := readInput(6, 1, test)
	score := d6p1.TotalScore(lines)
	fmt.Println(score)
}

func day6_part2(test bool) {
	lines := readInput(6, 2, test)
	score := d6p2.TotalScore(lines)
	fmt.Println(score)
}

func day7_part1(test bool) {
	lines := readInput(7, 1, test)
	score := d7p1.TotalScore(lines)
	fmt.Println(score)
}
