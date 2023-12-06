package d1p2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtract(t *testing.T) {
	assert.Equal(t, 29, ExtractCalibrationValue("two1nine"))
	assert.Equal(t, 83, ExtractCalibrationValue("eightwothree"))
	assert.Equal(t, 13, ExtractCalibrationValue("abcone2threexyz"))
	assert.Equal(t, 24, ExtractCalibrationValue("xtwone3four"))
	assert.Equal(t, 42, ExtractCalibrationValue("4nineeightseven2"))
	assert.Equal(t, 14, ExtractCalibrationValue("zoneight234"))
	assert.Equal(t, 76, ExtractCalibrationValue("7pqrstsixteen"))
	assert.Equal(t, 83, ExtractCalibrationValue("837ninethreezdcdbmjtph"))
}

func TestScore(t *testing.T) {
	assert.Equal(
		t,
		281,
		TotalScore([]string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}),
	)
}
