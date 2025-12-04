package d3p2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMaximumJoltage(t *testing.T) {
	N := 12
	assert.Equal(t, 434234234278, MaximumJoltage("234234234234278", N))
	assert.Equal(t, 987654321111, MaximumJoltage("987654321111111", N))
	assert.Equal(t, 811111111119, MaximumJoltage("811111111111119", N))
	assert.Equal(t, 434234234278, MaximumJoltage("234234234234278", N))
	assert.Equal(t, 888911112111, MaximumJoltage("818181911112111", N))
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		3121910778619,
		TotalScore([]string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}),
	)

}
