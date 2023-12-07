package d3p1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateBoundary(t *testing.T) {
	assert.Equal(t, "&$%%%%****", CalculateBoundary([]string{"%%%%", "&12$", "****"}, 0, 3, 1))
	assert.Equal(t, "$%%%***", CalculateBoundary([]string{"%%%", "12$", "***"}, -1, 2, 1))
	assert.Equal(t, "&%%%***", CalculateBoundary([]string{"%%%", "&12", "***"}, 0, 3, 1))
}

func TestGoodPartNums(t *testing.T) {
	assert.Equal(t, []int{12}, GetGoodPartNums([]string{
		"....",
		"*12.",
		"....",
	}))
	assert.Equal(t, []int{12, 34}, GetGoodPartNums([]string{
		"$.......",
		".12..34+",
		"........",
	}))
	assert.Equal(t, []int{467, 35}, GetGoodPartNums([]string{
		"467..114..",
		"...*......",
		"..35..633.",
	}))
	assert.Equal(t, []int{467, 35, 633, 617, 592, 755, 664, 598}, GetGoodPartNums([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}))
}

func TestScore(t *testing.T) {
	assert.Equal(t, 4361, TotalScore([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}))
}
