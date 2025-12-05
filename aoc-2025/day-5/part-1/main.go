package d5p1

import (
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func TotalScore(lines []string) int {
	score := 0

	ingredients, fresh_ids_str := GetIngredientsAndRanges(lines)

	fresh_ids := []Range{}
	for _, id_range := range fresh_ids_str {
		id_range_list := strings.Split(id_range, "-")
		min, _ := strconv.Atoi(id_range_list[0])
		max, _ := strconv.Atoi(id_range_list[1])
		fresh_ids = append(fresh_ids, Range{Min: min, Max: max})
	}

	for _, id_str := range ingredients {
		id, _ := strconv.Atoi(id_str)
		for _, id_range := range fresh_ids {
			if id >= id_range.Min && id <= id_range.Max {
				score++
				break
			}
		}
	}

	return score
}

func GetIngredientsAndRanges(input []string) ([]string, []string) {
	at_divider := false
	ranges := []string{}
	ingredients := []string{}

	for _, line := range input {
		if line == "" {
			at_divider = true
			continue
		}
		if !at_divider {
			ranges = append(ranges, line)
		} else {
			ingredients = append(ingredients, line)
		}
	}

	return ingredients, ranges
}
