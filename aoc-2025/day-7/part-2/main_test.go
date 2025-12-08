package d7p2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToList(t *testing.T) {
	t.Skip()
	assert.Equal(
		t,
		[]X{
			{index: 10, count: 2},
		},
		AddToList([]X{}, 10, 2),
	)

	assert.Equal(
		t,
		[]X{
			{index: 10, count: 10},
		},
		AddToList([]X{
			{index: 10, count: 9},
		}, 10, 1),
	)

	assert.Equal(
		t,
		[]X{
			{index: 2, count: 1},
			{index: 10, count: 1},
		},
		AddToList([]X{
			{index: 2, count: 1},
		}, 10, 1),
	)
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		4,
		TotalScore([]string{
			".......S.......",
			".......|.......",
			".......^.......",
			"......|.|......",
			"......^.^......",
			".....|.|.|.....",
		}),
	)

	t.SkipNow()
	assert.Equal(
		t,
		40,
		TotalScore([]string{
			".......S.......",
			".......|.......",
			".......^.......",
			"......|.|......",
			"......^.^......",
			".....|.|.|.....",
			".....^.^.^.....",
			"....|.|.|.|....",
			"....^.^.|.^....",
			"...|.|.|||.|...",
			"...^.^.||^.^...",
			"..|.|.|||.|.|..",
			"..^.|.^||.|.^..",
			".|.|||.||..|.|.",
			".^.^|^.^|^.|.^.",
			"|.|.|.|.|.|||.|",
		}),
	)
}
