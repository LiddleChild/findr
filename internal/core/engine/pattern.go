package engine

import "strings"

type Pattern struct {
	pattern string
	lps     []int
}

func CreatePattern(pattern string) *Pattern {
	pattern = strings.ToLower(pattern)
	lps := make([]int, len(pattern))
	i, j := 1, 0

	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return &Pattern{
		pattern,
		lps,
	}
}

func (p *Pattern) Match(content string) ([]int, bool) {
	content = strings.ToLower(content)
	index := make([]int, 0)
	if len(p.lps) == 0 {
		return index, true
	}

	contentLen := len(content)
	pattern, lps := p.pattern, p.lps

	i, j := 0, 0
	for i < contentLen {
		if pattern[j] == content[i] {
			i++
			j++
		}

		if j == len(pattern) {
			index = append(index, i-j)
			j = lps[j-1]
		} else if i < contentLen && pattern[j] != content[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}

	if len(index) == 0 {
		return index, false
	}

	return index, true
}
