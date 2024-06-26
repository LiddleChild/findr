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
	if !arg.CaseSensitive {
		path = strings.ToLower(path)
	}

	indices := pattern.Match(path)
	if len(indices) > 0 {
		highlighted := utils.HighlightByIndexes(path, indices, pattern.Len(), color.FgRed)
		fmt.Printf("%s\n", strings.TrimSpace(highlighted))
	}

	return nil
}

func handleContentSearch(path string, arg *Argument) errorwrapper.ErrorWrapper {
	f, err := os.Open(path)
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Core,
			err,
			"error occured while opening file")
	}

	snippets, err := pattern.MatchFile(f, arg.CaseSensitive)
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Core,
			err,
			"error occured while reading file")
	} else if len(snippets) > 0 {
		color.New(color.FgHiBlack).Println(path)
		for _, snip := range snippets {
			highlighted := utils.HighlightByIndexes(snip.Text, snip.Col, pattern.Len(), color.FgRed)
			fmt.Printf("Ln %d: %s\n", snip.Line, strings.TrimSpace(highlighted))
		}
		fmt.Println()
	}

	return nil
}
