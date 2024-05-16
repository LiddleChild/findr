package core

import (
	"os"
	"strings"
	"unicode"

	"github.com/LiddleChild/findr/utils"
)

type Pattern struct {
	pattern []rune
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
		[]rune(pattern),
		lps,
	}
}

func (p Pattern) Match(content string) []int {
	contentRunes := []rune(strings.ToLower(content))
	indices := make([]int, 0)

	i, j := 0, 0
	for i < len(content) {
		if contentRunes[i] == p.pattern[j] {
			if j == p.Len()-1 {
				indices = append(indices, i-j)
				j = p.lps[j-1]
			} else {
				i++
				j++
			}
		} else {
			if j > 0 {
				j = p.lps[j-1]
			} else if i < len(content) {
				i++
			}
		}
	}

	return indices
}

func (p Pattern) MatchFile(file *os.File) ([]Snippet, error) {
	snippets := make([]Snippet, 0)

	reader := utils.NewReader(file)

	ln := 1
	var sb strings.Builder
	newlineIndex := -1
	needResolve := false

	i, j := 0, 0
	r := reader.NextRune()
	sb.WriteRune(r)

	next := func() {
		r = reader.NextRune()
		sb.WriteRune(r)
		i++

		if r == '\n' {
			if needResolve {
				utils.Last(snippets).Text = sb.String()
				needResolve = false
			}

			newlineIndex = i
			ln++
			sb.Reset()
		}
	}

	for !reader.IsEOF() {
		if unicode.ToLower(r) == p.pattern[j] {
			if j == p.Len()-1 {
				if len(snippets) == 0 || !needResolve {
					snippets = append(snippets, Snippet{
						Line: ln,
						Col:  []int{i - j - newlineIndex - 1},
					})
				} else {
					last := utils.Last(snippets)
					last.Col = append(last.Col, i-j-newlineIndex-1)
				}

				needResolve = true
				j = p.lps[j-1]
			} else {
				next()
				j++
			}
		} else {
			if j > 0 {
				j = p.lps[j-1]
			} else if !reader.IsEOF() {
				next()
			}
		}
	}

	return snippets, nil
}

func (p Pattern) Len() int {
	return len(p.pattern)
}
