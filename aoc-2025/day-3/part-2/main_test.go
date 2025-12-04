package d3p2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolNew(t *testing.T) {
	assert.Equal(t, 434234234278, SolNew("234234234234278"))
	t.SkipNow()
	assert.Equal(t, 987654321111, SolNew("987654321111111"))
	assert.Equal(t, 811111111119, SolNew("811111111111119"))
	assert.Equal(t, 434234234278, SolNew("234234234234278"))
	assert.Equal(t, 888911112111, SolNew("818181911112111"))
}


func TestSol2(t *testing.T) {
	t.Skip()
	x := Sol2([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 12)
	assert.Equal(t, x.Value, 8)
	assert.Equal(t, x.Index, 0)

	y := Sol2([]int{1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 11)
	assert.Equal(t, y.Value, 8)
	assert.Equal(t, y.Index, 1)

	z := Sol2([]int{1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 11)
	assert.Equal(t, z.Value, 8)
	assert.Equal(t, z.Index, 1)
}

func TestMaxInList(t *testing.T) {
	t.Skip()
	assert.Equal(t, []int{1}, PopMax([]int{1, 8}))
	assert.Equal(t, []int{1, 2}, PopMax([]int{1, 2, 8}))
	assert.Equal(t, []int{1, 2, 3}, PopMax([]int{1, 2, 8, 3}))
	assert.Equal(
		t,
		[]int{8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
		PopMax([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}),
	)
}

func TestMaximumJoltage(t *testing.T) {
	t.SkipNow()
	assert.Equal(t, 987654321111, MaximumJoltage("987654321111111"))
	assert.Equal(t, 811111111119, MaximumJoltage("811111111111119"))
	assert.Equal(t, 434234234278, MaximumJoltage("234234234234278"))
	assert.Equal(t, 888911112111, MaximumJoltage("818181911112111"))
}

func TestScore(t *testing.T) {
	t.Skip()
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
