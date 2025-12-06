package d6p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		4277556,
		TotalScore([]string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
			"*   +   *   +  ",
		}),
	)
}
