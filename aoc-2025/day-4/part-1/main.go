package d4p1

type GridSize struct {
	Rows int
	Cols int
}

type Position struct {
	i int
	j int
}

var neighborOffsets = []Position{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func TotalScore(grid []string) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	score := 0
	gs := GridSize{Rows: len(grid), Cols: len(grid[0])}

	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '@' {
				continue
			}

			count := 0
			for _, offset := range neighborOffsets {
				neighborPos := Position{i: i + offset.i, j: j + offset.j}
				if CheckIfRoll(grid, gs, neighborPos) {
					count++
				}
			}

			if count < 4 {
				score++
			}
		}
	}
	return score
}

func CheckIfRoll(grid []string, grid_size GridSize, pos Position) bool {
	if pos.i < 0 || pos.j < 0 || pos.i >= grid_size.Rows || pos.j >= grid_size.Cols {
		return false
	}
	return grid[pos.i][pos.j] == '@'
}
