package d3lib

// CalculateBoundary returns the left, right and top+bottom lines of strings
// that bounds a given substr inside the input s. This substr is defined by
// left, right and center.
//
// left is the left index (exclusive) of chars to calculate boundary of.
// right is right index (exclusive) of chars to calculate boundary of.
// center is index of row containing chars to calculate boundary of.
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
