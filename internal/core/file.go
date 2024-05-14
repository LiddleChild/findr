package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/utils"
	"github.com/fatih/color"
)

type dirNode struct {
	path  string
	depth int
}

var pattern *Pattern

func HandleFilenameSearch(path string, arg *Argument) errorwrapper.ErrorWrapper {
	idx, ok := pattern.Match(path)
	if ok {
		highlighted := utils.HighlightByIndexes(path, idx, pattern.Len(), color.FgRed)
		fmt.Printf("%s\n\n", strings.TrimSpace(highlighted))
	}

	return nil
}

func handleContentSearch(path string, _ *Argument) errorwrapper.ErrorWrapper {
	bs, err := os.ReadFile(path)
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Core,
			err,
			"error occured while reading files")
	}
	content := string(bs)

	index, ok := pattern.Match(content)
	if ok {
		mln := utils.ToMultiline(content)

		underline := color.New(color.Underline).SprintFunc()

		fmt.Println(underline(path))
		for _, i := range index {
			ln, col, s := mln.GetSnippet(i)
			highlighted := utils.HighlightByIndexes(s, []int{col}, pattern.Len(), color.FgRed)
			fmt.Printf("Ln %d, Col %d: %s\n", ln, col, strings.TrimSpace(highlighted))
		}
		fmt.Println()
	}

	return nil
}

func Traverse(arg *Argument) errorwrapper.ErrorWrapper {
	st := utils.NewStack[dirNode]()
	st.Push(dirNode{
		path:  arg.WorkingDirectory,
		depth: 0,
	})

	pattern = CreatePattern(arg.Query)

	for st.Size() > 0 {
		dir := st.Top()
		st.Pop()

		entries, err := os.ReadDir(dir.path)
		if err != nil {
			return errorwrapper.NewWithMessage(
				errorwrapper.Core,
				err,
				"error occured while reading directories")
		}

		for _, e := range entries {
			if _, ok := arg.IgnoredPaths[e.Name()]; ok {
				continue
			}

			path := filepath.Join(dir.path, e.Name())
			if e.IsDir() && dir.depth < arg.MaxDepth {
				st.Push(dirNode{
					path,
					dir.depth + 1,
				})
			}

			var werr errorwrapper.ErrorWrapper
			if arg.ContentSearch && !e.IsDir() {
				werr = handleContentSearch(path, arg)
			} else {
				werr = HandleFilenameSearch(path, arg)
			}

			if werr != nil {
				return werr
			}
		}
	}

	return nil
}
