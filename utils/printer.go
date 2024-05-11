package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func HighlightedPrintln(text string, pattern string, c color.Attribute) {
	paint := color.New(c).SprintFunc()
	fmt.Println(strings.ReplaceAll(text, pattern, paint(pattern)))
}
