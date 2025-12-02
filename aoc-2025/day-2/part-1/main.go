package d2p1

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

func NumIsInvalid(num string) bool {
	return num[0] == '0' || (len(num)%2 == 0 && num[:len(num)/2] == num[len(num)/2:])
}
