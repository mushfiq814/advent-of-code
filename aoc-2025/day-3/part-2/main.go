package d3p2

import (
	"fmt"
	"math"
	"sort"
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
		score += MaximumJoltage(line)
	}

	return score
}

func MaximumJoltage(bank string) int {
	BATT_COUNT := 12

	// get string array
	batteries_str := strings.Split(bank, "")

	// build int array
	batteries := []int{}
	for _, battery := range batteries_str {
		b, _ := strconv.Atoi(battery)
		batteries = append(batteries, b)
	}

	// store value with index
	indexed := make([]IndexedValue, len(batteries))
	for i, v := range batteries {
		indexed[i] = IndexedValue{Index: i, Value: v}
	}

	// reverse sort with index while keeping ascending order of index for the
	// same value
	sort.Slice(indexed, func(i, j int) bool {
		if indexed[i].Value != indexed[j].Value {
			return indexed[i].Value > indexed[j].Value
		}
		return indexed[i].Index < indexed[j].Index
	})

	// if highest number is at the end, it cannot be considered the first digit
	start := 0
	if indexed[start].Index >= len(indexed)-BATT_COUNT+1 {
		start++
	}

	// get the first digit
	d0 := indexed[start]

	// remove first digit from indexed
	indexed_without_d0 := []IndexedValue{}
	if start == 0 {
		indexed_without_d0 = append(indexed_without_d0, indexed[1:][:]...)
	} else {
		indexed_without_d0 = append(indexed_without_d0, indexed[:start][:]...)
		indexed_without_d0 = append(indexed_without_d0, indexed[start+1:][:]...)
	}

	// indexed_without_d0 [{8 1} {7 2} {6 3} {5 4} {4 5} {3 6} {2 7} {1 8} {1 9} {1 10} {1 11} {1 12} {1 13} {1 14}]

	for i := 0; i < BATT_COUNT; i++ {
		for j := 0; j < len(indexed_without_d0)-1; j++ {
			if indexed_without_d0[j].Index > d0.Index {
				return d0.Value*10 + indexed_without_d0[j].Value
			}
		}
	}

	return 0
}

func PopMax(list []int) []int {
	max := 0
	max_i := 0
	for i, item := range list {
		if item > max {
			max = item
			max_i = i
		}
	}

	result := []int{}
	if max_i == 0 {
		result = append(result, list[1:][:]...)
	} else {
		result = append(result, list[:max_i][:]...)
		result = append(result, list[max_i+1:][:]...)
	}

	return result
}

func SolNew(bank string) int {
	N := 12

	// get string array
	batteries_str := strings.Split(bank, "")

	// build int array
	batteries := []int{}
	for _, battery := range batteries_str {
		b, _ := strconv.Atoi(battery)
		batteries = append(batteries, b)
	}

	z := []int{}
	for N > 0 {
		x := Sol2(batteries, N)
		batteries = RemoveItemFromArray(batteries, x.Index)
		z = append(z, x.Value)
		N--
	}
	fmt.Println(z)

	return CombineIntoDigits(z)
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
	if index == 0 {
		result = append(result, arr[1:][:]...)
	} else {
		result = append(result, arr[:index][:]...)
		result = append(result, arr[index+1:][:]...)
	}
	return result
}

func Sol2(list []int, N int) IndexedValue {
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
