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

func Traverse(query string, arg *Argument) errorwrapper.ErrorWrapper {
	st := utils.NewStack[dirNode]()
	st.Push(dirNode{
		path:  arg.WorkingDirectory,
		depth: 0,
	})

	pattern := CreatePattern(query)

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
			path := filepath.Join(dir.path, e.Name())

			if arg.ContentSearch && !e.IsDir() {
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

					red := color.New(color.FgRed).SprintFunc()
					underline := color.New(color.Underline).SprintFunc()

					fmt.Println(underline(path))
					for _, i := range index {
						ln, col, s := mln.GetSnippet(i)
						fmt.Printf("Ln %d, Col %d: %s\n", ln, col, strings.TrimSpace(strings.ReplaceAll(s, query, red(query))))
					}
				}
			} else {
				red := color.New(color.FgRed).SprintFunc()
				if strings.Contains(path, query) {
					fmt.Println(strings.TrimSpace(strings.ReplaceAll(path, query, red(query))))
				}
			}

			if e.IsDir() && dir.depth < arg.MaxDepth {
				st.Push(dirNode{
					path,
					dir.depth + 1,
				})
			}
		}
	}

	return nil
}
