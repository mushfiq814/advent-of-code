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

	points := []Point{}
	for _, line := range lines {
		points = append(points, ParseCoordsToPoints(line))
	}

	distance_grid := MakeDistanceGrid(len(lines))
	for i := range lines {
		for j := range lines {
			if i == j {
				distance_grid[i][j] = 0
				continue
			}
			// TODO: memoize + use matrix symmetry
			distance_grid[i][j] = Distance(points[i], points[j])
		}
	}

	fmt.Println(distance_grid)

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

func MakeDistanceGrid(n int) [][]float64 {
	distance_grid := make([][]float64, n)
	rows := make([]float64, n*n)
	for i := range n {
		distance_grid[i] = rows[i*n : (i+1)*n]
	}
	return distance_grid
}
