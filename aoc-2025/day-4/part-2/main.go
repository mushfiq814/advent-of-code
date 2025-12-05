package d4p2

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
	score := 0
	gs := GridSize{Rows: len(grid), Cols: len(grid[0])}

	next_score := 0
	grid, next_score = SubLoop(grid, gs)
	score += next_score
	for next_score > 0 {
		grid, next_score = SubLoop(grid, gs)
		score += next_score
	}

	return score
}

func SubLoop(grid []string, gs GridSize) ([]string, int) {
	score := 0
	new_grid := []string{}
	new_grid = append(new_grid, grid...)
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
				x := []rune(new_grid[i])
				x[j] = 'x'
				new_grid[i] = string(x)
			}
		}
	}
	return new_grid, score
}

func CheckIfRoll(grid []string, grid_size GridSize, pos Position) bool {
	if pos.i < 0 || pos.j < 0 || pos.i >= grid_size.Rows || pos.j >= grid_size.Cols {
		return false
	}
	return grid[pos.i][pos.j] == '@'
}
