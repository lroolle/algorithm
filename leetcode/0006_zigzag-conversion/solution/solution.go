package solution

import (
	"strings"
)

func convert(s string, numRows int) string {
	slines := make([][]rune, numRows)
	srunes := []rune(s)
	n := numRows + (numRows - 2)
	if n <= 1 {
		return s
	}

	for i, c := range srunes {
		lineNum := i % n
		if lineNum >= numRows {
			lineNum = (numRows - 1) - (lineNum % (numRows - 1))
		}
		slines[lineNum] = append(slines[lineNum], c)
	}
	ret := []string{}
	for _, sline := range slines {
		ret = append(ret, string(sline))
	}
	return strings.Join(ret, "")
}

func convert2(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	slines := make([]string, numRows)
	srunes := []rune(s)
	row, steps := 0, 1

	for _, c := range srunes {
		slines[row] += string(c)
		if row == 0 {
			steps = 1
		} else if row == numRows-1 {
			steps = -1
		}
		row += steps
	}
	return strings.Join(slines, "")
}
