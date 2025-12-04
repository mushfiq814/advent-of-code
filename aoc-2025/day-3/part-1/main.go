package d3p1

import (
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
	if indexed[start].Index == len(indexed)-1 {
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

	for i := 0; i < len(indexed_without_d0)-1; i++ {
		if indexed_without_d0[i].Index > d0.Index {
			return d0.Value*10 + indexed_without_d0[i].Value
		}
	}

	return 0
}
