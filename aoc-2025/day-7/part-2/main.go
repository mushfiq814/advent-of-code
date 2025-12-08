package d7p2

import (
	"fmt"
	"slices"
)

type X struct {
	index int
	count int
}

func TotalScore(lines []string) int {
	score := 1
	beam_counts := []X{}

	beam_indices := []int{}
	for i, char := range lines[0] {
		if char == 'S' {
			beam_indices = []int{i}
			beam_counts = []X{{index: i, count: 1}}
			break
		}
	}

	for i := 2; i < len(lines); i++ {
		new_beam_indices := []int{}
		new_beam_counts := []X{}

		for _, beam_index := range beam_indices {
			existing_count := FindXFromIndex(beam_counts, beam_index).count
			if lines[i][beam_index] == '^' {
				if beam_index-1 >= 0 && beam_index-1 < len(lines[i]) && !slices.Contains(new_beam_indices, beam_index-1) {
					new_beam_indices = append(new_beam_indices, beam_index-1)
					new_beam_counts = AddToList(new_beam_counts, beam_index-1, existing_count)
				}
				if beam_index+1 >= 0 && beam_index+1 < len(lines[i]) && !slices.Contains(new_beam_indices, beam_index+1) {
					new_beam_indices = append(new_beam_indices, beam_index+1)
					new_beam_counts = AddToList(new_beam_counts, beam_index+1, existing_count)
				}
			} else {
				if !slices.Contains(new_beam_indices, beam_index) {
					new_beam_indices = append(new_beam_indices, beam_index)
					new_beam_counts = AddToList(new_beam_counts, beam_index, existing_count)
				}
			}
		}

		beam_indices = new_beam_indices
		beam_counts = new_beam_counts
		fmt.Println(beam_counts)
	}

	return score
}

func FindXFromIndex(list []X, index int) X {
	for i := range list {
		if list[i].index == index {
			return list[i]
		}
	}
	return X{}
}

func AddToList(list []X, index int, count int) []X {
	found := false
	for i := range list {
		if list[i].index == index {
			found = true
			list[i].count += count
			break
		}
	}
	if !found {
		list = append(list, X{index: index, count: count})
	}
	return list
}
