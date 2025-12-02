package d2p2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorsOf(t *testing.T) {
	t.Skip()
	assert.Equal(t, []int{1}, FactorsOf(1))
	assert.Equal(t, []int{1, 2}, FactorsOf(2))
	assert.Equal(t, []int{1, 2, 4}, FactorsOf(4))
	assert.Equal(t, []int{1, 5}, FactorsOf(5))
	assert.Equal(t, []int{1, 2, 5, 10}, FactorsOf(10))
}

func TestInvalid(t *testing.T) {
	assert.Equal(t, false, NumIsInvalid("2121212118"))
	assert.Equal(t, true, NumIsInvalid("121212"))
	assert.Equal(t, false, NumIsInvalid("123"))
	assert.Equal(t, true, NumIsInvalid("111"))
	assert.Equal(t, true, NumIsInvalid("55"))
	assert.Equal(t, true, NumIsInvalid("6464"))
	assert.Equal(t, true, NumIsInvalid("6464"))
	assert.Equal(t, true, NumIsInvalid("666"))
	assert.Equal(t, true, NumIsInvalid("565656"))
	assert.Equal(t, true, NumIsInvalid("824824824"))
	assert.Equal(t, true, NumIsInvalid("2121212121"))
	assert.Equal(t, true, NumIsInvalid("6666"))
	assert.Equal(t, true, NumIsInvalid("123123"))
	assert.Equal(t, true, NumIsInvalid("1010"))
	assert.Equal(t, true, NumIsInvalid("0101"))
	assert.Equal(t, false, NumIsInvalid("101"))
	assert.Equal(t, true, NumIsInvalid("12121212"))
	assert.Equal(t, true, NumIsInvalid("446446"))
	assert.Equal(t, true, NumIsInvalid("38593859"))
	assert.Equal(t, true, NumIsInvalid("1188511885"))
	assert.Equal(t, false, NumIsInvalid("2121212118"))
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		4174379265,
		TotalScore([]string{
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
		}),
	)
}
