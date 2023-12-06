package d1p1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtract(t *testing.T) {
	assert.Equal(t, 12, ExtractCalibrationValue("1abc2"))
	assert.Equal(t, 38, ExtractCalibrationValue("pqr3stu8vwx"))
	assert.Equal(t, 15, ExtractCalibrationValue("a1b2c3d4e5f"))
	assert.Equal(t, 77, ExtractCalibrationValue("treb7uchet"))
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		142,
		TotalScore([]string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}),
	)
}
