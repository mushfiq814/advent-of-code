package d1p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		11,
		TotalScore([]string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}),
	)
}
