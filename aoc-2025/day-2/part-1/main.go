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
					count+=i
				}
			}
		}
	}

	return count
}

// only twice repeated: 99 but NOT 999
// one num repeated: 11
// two nums repeated: 2020
// three nums repeated: 123123
// even numbered length with boundary halfway: 1212 has 4 digits with boundary before index 2
// no leading zeros: 01 or 001 or 0100
func NumIsInvalid(num string) bool {
	if num[0] == '0' {
		return true
	}

	if !(len(num)%2 == 0) {
		return false
	}

	if num[0:len(num)/2] == num[len(num)/2:] {
		return true
	}

	return false
}
