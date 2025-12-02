package d2p2

import (
	"strconv"
	"strings"
)

func TotalScore(lines []string) int {
	count := 0

	for _, line := range lines {
		groups := strings.Split(line, ",")
		for _, group := range groups {
			nums := strings.Split(group, "-")
			start_str := nums[0]
			end_str := nums[1]

			start, _ := strconv.Atoi(start_str)
			end, _ := strconv.Atoi(end_str)

			for i := start; i <= end; i++ {
				if NumIsInvalid(strconv.Itoa(i)) {
					count += i
				}
			}
		}
	}

	return count
}

func FactorsOf(num int) []int {
	if num <= 1 {
		return []int{1}
	}

	res := []int{1}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			res = append(res, i)
		}
	}

	res = append(res, num)

	return res
}

func NumIsInvalid(num string) bool {
	match := false

	for _, i := range FactorsOf(len(num)) {
		for j := 0; j < len(num)-i; j += i {
			match = num[j:j+i] == num[j+i:j+2*i]
			if !match {
				break
			}
		}
		if match {
			break
		}
	}

	return match
}
