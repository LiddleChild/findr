package utils

import (
	"strings"

	"github.com/fatih/color"
)

func HighlightByIndexes(content string, index []int, size int, attr color.Attribute) string {
	if len(index) == 0 {
		return content
	}

	highlighter := color.New(attr).SprintFunc()

	var builder strings.Builder
	lastIndex := 0
	for i, idx := range index {
		builder.WriteString(content[lastIndex:idx])
		builder.WriteString(highlighter(content[idx : idx+size]))
		lastIndex = i + size
	}

	builder.WriteString(content[index[len(index)-1]+size:])

	return builder.String()
}
