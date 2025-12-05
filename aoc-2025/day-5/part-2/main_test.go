package d5p2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		6,
		TotalScore([]string{
			"1-3",
			"10-12",
			"",
		}),
	)

	assert.Equal(
		t,
		8,
		TotalScore([]string{
			"5-10",
			"8-12",
			"",
		}),
	)

	assert.Equal(
		t,
		13,
		TotalScore([]string{
			"10-20",
			"12-15",
			"21-22",
			"",
		}),
	)

	assert.Equal(
		t,
		38,
		TotalScore([]string{
			"50-60",
			"10-20",
			"15-25",
			"70-80",
			"75-75",
			"",
		}),
	)

	assert.Equal(
		t,
		14,
		TotalScore([]string{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
			"",
			"1",
			"5",
			"8",
			"11",
			"17",
			"32",
		}),
	)
}
