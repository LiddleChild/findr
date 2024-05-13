package utils

import (
	"strings"
)

type Multiline struct {
	lines []string
	lenx  []int
}

func ToMultiline(text string) Multiline {
	lines := strings.Split(text, "\n")
	lenx := make([]int, len(lines)+1)

	for i, line := range lines {
		lenx[i+1] = lenx[i] + len(line) + 1
	}

	return Multiline{
		lines,
		lenx,
	}
}

func (mln Multiline) GetSnippet(index int) (int, int, string) {
	i := 1
	for index > mln.lenx[i] && i < len(mln.lenx)-1 {
		i++
	}

	return i, index - mln.lenx[i-1], mln.lines[i-1]
}
