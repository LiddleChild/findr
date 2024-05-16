package utils

import (
	"strings"

	"github.com/fatih/color"
)

func HighlightByIndexes(content string, indices []int, size int, attr color.Attribute) string {
	if len(indices) == 0 {
		return content
	}

	rs := []rune(content)

	highlighter := color.New(attr).SprintFunc()

	var builder strings.Builder
	lastIndex := 0
	for _, index := range indices {
		writeRunes(&builder, rs[lastIndex:index])
		builder.WriteString(highlighter(string(rs[index : index+size])))
		lastIndex = index + size
	}

	writeRunes(&builder, rs[indices[len(indices)-1]+size:])

	return builder.String()
}

func writeRunes(builder *strings.Builder, rs []rune) {
	for _, r := range rs {
		builder.WriteRune(r)
	}
}
