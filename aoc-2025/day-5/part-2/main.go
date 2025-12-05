package d5p2

import (
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

// 346934484731209 too low
func TotalScore(lines []string) int {
	score := 0
	ids := FormatRanges(GetRanges(lines))

	sort.Slice(ids, func(i, j int) bool {
		if ids[i].Min == ids[j].Min {
			return ids[i].Max < ids[j].Max
		}
		return ids[i].Min < ids[j].Min
	})

	// start with first item in range
	ids2 := []Range{}
	ids2 = append(ids2, ids[0])

	// look for overlapping ranges using 2 markers for each array
	i_ids2 := 0
	i_ids := 1
	for i_ids < len(ids) {
		if ids[i_ids].Min <= ids2[i_ids2].Max {
			// if there is an overlap, replace the current item with the broader range
			new_min := ids2[i_ids2].Min
			new_max := Max(ids[i_ids].Max, ids2[i_ids2].Max)
			ids2[i_ids2] = Range{Min: new_min, Max: new_max}
			// _only_ advance ids marker
			i_ids++
		} else {
			// if there is no overlap, add the current item
			ids2 = append(ids2, ids[i_ids])
			// advance both markers
			i_ids++
			i_ids2++
		}
	}

	for _, rng := range ids2 {
		score += rng.Max - rng.Min + 1
	}

	return score
}

func GetRanges(input []string) []string {
	ranges := []string{}

	for _, line := range input {
		if line == "" {
			break
		}
		ranges = append(ranges, line)
	}

	return ranges
}

func FormatRanges(fresh_ids_str []string) []Range {
	fresh_ids := []Range{}
	for _, id_range := range fresh_ids_str {
		id_range_list := strings.Split(id_range, "-")
		min, _ := strconv.Atoi(id_range_list[0])
		max, _ := strconv.Atoi(id_range_list[1])
		fresh_ids = append(fresh_ids, Range{Min: min, Max: max})
	}
	return fresh_ids
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
