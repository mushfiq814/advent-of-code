package d1p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		0,
		TotalScore([]string{
			"1",
			"2",
		}),
	)

}
