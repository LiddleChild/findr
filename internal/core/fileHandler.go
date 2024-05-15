package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/utils"
	"github.com/fatih/color"
)

func HandleFilenameSearch(path string, arg *Argument) errorwrapper.ErrorWrapper {
	idx, ok := pattern.Match(path)
	if ok {
		highlighted := utils.HighlightByIndexes(path, idx, pattern.Len(), color.FgRed)
		fmt.Printf("%s\n", strings.TrimSpace(highlighted))
	}

	return nil
}

func handleContentSearch(path string, _ *Argument) errorwrapper.ErrorWrapper {
	f, err := os.Open(path)
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Core,
			err,
			"error occured while opening file")
	}

	snippets, ok, err := pattern.MatchFile(f)
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Core,
			err,
			"error occured while reading file")
	} else if ok {
		fmt.Println(path)
		for _, snip := range snippets {
			highlighted := utils.HighlightByIndexes(snip.Text, snip.Col, pattern.Len(), color.FgRed)
			fmt.Printf("Ln %d: %s\n", snip.Line, strings.TrimSpace(highlighted))
		}
		fmt.Println()
	}

	return nil
}
