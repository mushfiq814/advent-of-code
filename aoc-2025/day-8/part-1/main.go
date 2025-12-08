package d8p1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

func TotalScore(lines []string) int {
	score := 0
	fmt.Println("TODO")

	for _, line := range lines {
		fmt.Println(line)
	}

	return score
}

func ParseCoordsToPoints(coord_str string) Point {
	coord_list := []int{}
	for c := range strings.SplitSeq(coord_str, ",") {
		n, _ := strconv.Atoi(c)
		coord_list = append(coord_list, n)
	}
	return Point{x: coord_list[0], y: coord_list[1], z: coord_list[2]}
}

func Distance(p1 Point, p2 Point) float64 {
	x := math.Pow(float64(p1.x-p2.x), 2)
	y := math.Pow(float64(p1.y-p2.y), 2)
	z := math.Pow(float64(p1.z-p2.z), 2)
	return math.Sqrt(x + y + z)
}
