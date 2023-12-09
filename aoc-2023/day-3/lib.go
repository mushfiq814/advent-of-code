package d3lib

func CalculateBoundary(
	s []string,
	left int,
	right int,
	center int,
) (
	string,
	string,
	string,
) {
	L := ""
	R := ""
	TB := ""

	if left >= 0 {
		L = s[center][left : left+1]
	}
	if right+1 <= len(s[center]) {
		R = s[center][right : right+1]
	}

	for _, r := range []int{center - 1, center + 1} {
		adjustedLeft := left
		adjustedRight := right
		if r >= 0 && r < len(s) {
			if left < 0 {
				adjustedLeft = 0
			}
			if right+1 > len(s[r]) {
				adjustedRight = len(s[r]) - 1
			}
			TB += s[r][adjustedLeft : adjustedRight+1]
		}
	}

	return L, R, TB
}
