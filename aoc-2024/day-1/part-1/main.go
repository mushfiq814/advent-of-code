package d1p1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func TotalScore(lines []string) int {
	totalScore := 0

	list1 := []int{}
	list2 := []int{}

	for _, line := range lines {
    parts := strings.Split(line, "   ")

    num1, err1 := strconv.Atoi(parts[0])
    num2, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil {
      panic(err1)
    }

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		totalScore += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return totalScore
}
