package d7p1

import (
	"slices"
)

func TotalScore(lines []string) int {
	score := 0

	beam_indices := []int{IndexOf(lines[0], 'S')}

	for i := 2; i < len(lines); i++ {
		new_beam_indicex := []int{}

		for cell_index, cell := range lines[i] {
			if slices.Contains(beam_indices, cell_index) {
				if cell == '^' {
					score++
					if cell_index-1 >= 0 && cell_index-1 < len(lines[i]) && !slices.Contains(new_beam_indicex, cell_index-1) {
						new_beam_indicex = append(new_beam_indicex, cell_index-1)
					}
					if cell_index+1 >= 0 && cell_index+1 < len(lines[i]) && !slices.Contains(new_beam_indicex, cell_index+1) {
						new_beam_indicex = append(new_beam_indicex, cell_index+1)
					}
				} else {
					if !slices.Contains(new_beam_indicex, cell_index) {
						new_beam_indicex = append(new_beam_indicex, cell_index)
					}
				}
			}
		}
		beam_indices = new_beam_indicex
	}

	return score
}

func IndexOf(str string, sub_str byte) int {
	for i := 0; i < len(str); i++ {
		if str[i] == sub_str {
			return i
		}
	}
	return -1
}
