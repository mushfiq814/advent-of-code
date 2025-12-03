package d3p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumJoltage(t *testing.T) {
	assert.Equal(t, 98, MaximumJoltage("987654321111111"))
	assert.Equal(t, 89, MaximumJoltage("811111111111119"))
	assert.Equal(t, 78, MaximumJoltage("234234234234278"))
	assert.Equal(t, 92, MaximumJoltage("818181911112111"))
}

func TestScore(t *testing.T) {
	t.Skip()
	assert.Equal(
		t,
		357,
		TotalScore([]string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}),
	)

}
