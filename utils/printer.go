package utils

import (
	"strings"

	"github.com/fatih/color"
)

func HighlightByIndexes(content string, index []int, size int, attr color.Attribute) string {
	if len(index) == 0 {
		return content
	}

	rs := []rune(content)

	highlighter := color.New(attr).SprintFunc()

	var builder strings.Builder
	lastIndex := 0
	for i, idx := range index {
		writeRunes(&builder, rs[lastIndex:idx])
		builder.WriteString(highlighter(string(rs[idx : idx+size])))
		lastIndex = i + size
	}

	writeRunes(&builder, rs[index[len(index)-1]+size:])

	return builder.String()
}

func writeRunes(builder *strings.Builder, rs []rune) {
	for _, r := range rs {
		builder.WriteRune(r)
	}
}
