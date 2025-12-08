package d8p1

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCoordsToPoints(t *testing.T) {
	assert.Equal(t, Point{x: 1, y: 2, z: 3}, ParseCoordsToPoints("1,2,3"))
	assert.Equal(t, Point{x: 162, y: 817, z: 812}, ParseCoordsToPoints("162,817,812"))
}

func TestDistance(t *testing.T) {
	assert.Equal(
		t,
		10.25,
		math.Round(
			Distance(
				ParseCoordsToPoints("7,4,3"),
				ParseCoordsToPoints("17,6,2"),
			)*100)/100)
}

func TestScore(t *testing.T) {
	t.Skip()
	assert.Equal(
		t,
		40,
		TotalScore([]string{
			"162,817,812",
			"57,618,57",
			"906,360,560",
			"592,479,940",
			"352,342,300",
			"466,668,158",
			"542,29,236",
			"431,825,988",
			"739,650,466",
			"52,470,668",
			"216,146,977",
			"819,987,18",
			"117,168,530",
			"805,96,715",
			"346,949,466",
			"970,615,88",
			"941,993,340",
			"862,61,35",
			"984,92,344",
			"425,690,689",
		}),
	)
}
