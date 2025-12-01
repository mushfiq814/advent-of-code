package d1p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.Equal(t, 0, MapToRealDial(400))
	assert.Equal(t, 98, MapToRealDial(-2))
	assert.Equal(t, 50, MapToRealDial(-150))
	assert.Equal(t, 0, MapToRealDial(0))
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		2,
		TotalScore([]string{
			"L50",
			"L1",
			"R1",
		}),
	)

	assert.Equal(
		t,
		1,
		TotalScore([]string{
			"L50",
			"R50",
		}),
	)

	assert.Equal(
		t,
		2,
		TotalScore([]string{
			"R20",
			"R20",
			"R10",
			"R50",
			"R50",
		}),
	)

	assert.Equal(
		t,
		3,
		TotalScore([]string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}),
	)
}
