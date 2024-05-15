package core

import (
	"os"
	"strings"
	"unicode"

	"github.com/LiddleChild/findr/utils"
)

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

func (p *Pattern) MatchFile(file *os.File) ([]Snippet, bool, error) {
	if len(p.lps) == 0 {
		return nil, true, nil
	}

	snippets := make([]Snippet, 0)

	pattern, lps := []rune(p.pattern), p.lps

	br := utils.NewReader(file)
	r, i, j := unicode.ToLower(br.NextRune()), 0, 0

	var sb strings.Builder
	indexLF := -1
	line := 1
	waiting := false // waiting for new line to finish snippet

	completeLine := func() {
		if r == '\n' {
			if waiting {
				snippets[len(snippets)-1].Text = sb.String()
				for i := range len(snippets[len(snippets)-1].Col) {
					snippets[len(snippets)-1].Col[i] -= indexLF + 1
				}
			}

			waiting = false
			sb.Reset()
			indexLF = i
			line++
		}
	}

	for !br.IsEOF() {
		if br.Error() != nil {
			return nil, false, br.Error()
		}

		if pattern[j] == r {
			r = unicode.ToLower(br.NextRune())
			i++
			if r != '\n' && r != '\r' {
				sb.WriteRune(r)
			}

			j++
		}

		if j == len(pattern) {
			if len(snippets) > 0 && snippets[len(snippets)-1].Line == line {
				snippets[len(snippets)-1].Col = append(
					snippets[len(snippets)-1].Col,
					i-j)

			} else {
				snippets = append(
					snippets,
					Snippet{
						Col:  []int{i - j},
						Line: line,
					})
			}

			waiting = true
			completeLine()

			j = lps[j-1]
		} else if !br.IsEOF() && pattern[j] != r {
			if j != 0 {
				j = lps[j-1]
			} else {
				r = unicode.ToLower(br.NextRune())
				i++
				if r != '\n' && r != '\r' {
					sb.WriteRune(r)
				}

				completeLine()
			}
		}
	}

	if len(snippets) == 0 {
		return snippets, false, nil
	}

	return snippets, true, nil
}

func (p *Pattern) Len() int {
	return len(p.pattern)
}
