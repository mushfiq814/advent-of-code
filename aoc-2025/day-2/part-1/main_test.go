package d2p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalid(t *testing.T) {
	assert.Equal(t, false, NumIsInvalid("123"))
	assert.Equal(t, true, NumIsInvalid("55"))
	assert.Equal(t, true, NumIsInvalid("6464"))
	assert.Equal(t, true, NumIsInvalid("6464"))
	assert.Equal(t, false, NumIsInvalid("666"))
	assert.Equal(t, true, NumIsInvalid("6666"))
	assert.Equal(t, true, NumIsInvalid("123123"))
	assert.Equal(t, true, NumIsInvalid("1010"))
	assert.Equal(t, true, NumIsInvalid("0101"))
	assert.Equal(t, false, NumIsInvalid("101"))
	assert.Equal(t, true, NumIsInvalid("12121212"))
	assert.Equal(t, true, NumIsInvalid("446446"))
	assert.Equal(t, true, NumIsInvalid("38593859"))
	assert.Equal(t, true, NumIsInvalid("1188511885"))
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		33,
		TotalScore([]string{
			"11-22",
		}),
	)

	assert.Equal(
		t,
		1188513027,
		TotalScore([]string{
			"11-22,95-115,998-1012,1188511880-1188511890",
		}),
	)

	assert.Equal(
		t,
		1227775554,
		TotalScore([]string{
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
		}),
	)
}
