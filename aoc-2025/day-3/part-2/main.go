package d3p2

import (
	"math"
	"strconv"
	"strings"
)

type IndexedValue struct {
	Value int
	Index int
}

func TotalScore(lines []string) int {
	score := 0

	for _, line := range lines {
		score += MaximumJoltage(line, 12)
	}

	return score
}

func MaximumJoltage(bank string, batt_count int) int {
	// get string array
	batteries_str := strings.Split(bank, "")

	// build int array
	batteries := []int{}
	for _, battery := range batteries_str {
		b, _ := strconv.Atoi(battery)
		batteries = append(batteries, b)
	}

	z := []int{}
	for batt_count > 0 {
		x := GetFirstDigit(batteries, batt_count)
		batteries = RemoveItemFromArray(batteries, x.Index)
		z = append(z, x.Value)
		batt_count--
	}

	return CombineIntoDigits(z)
}

func GetFirstDigit(list []int, N int) IndexedValue {
	stop_before := len(list) - N + 1

	max := 0
	max_i := 0
	for i := range stop_before {
		if list[i] > max {
			max = list[i]
			max_i = i
		}
	}

	return IndexedValue{Value: max, Index: max_i}
}

func CombineIntoDigits(list []int) int {
	result := 0
	for i := range list {
		result += list[len(list)-i-1] * int(math.Pow(10, float64(i)))
	}
	return result
}

func RemoveItemFromArray(arr []int, index int) []int {
	result := []int{}
	result = append(result, arr[index+1:][:]...)
	return result
}
