package d2p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPossibility(t *testing.T) {
	assert.Equal(t, true, getPossibility("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
	assert.Equal(t, true, getPossibility("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"))
	assert.Equal(t, false, getPossibility("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"))
	assert.Equal(t, false, getPossibility("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"))
	assert.Equal(t, true, getPossibility("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))

	// bad inputs: no game index
	assert.Equal(t, false, getPossibility("Game"))
	assert.Equal(t, false, getPossibility("Game : 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))
}

func getPossibility(game string) bool {
	possibility, _ := CalculateGamePossibility(game)
	return possibility
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		8,
		TotalScore([]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}),
	)
}
