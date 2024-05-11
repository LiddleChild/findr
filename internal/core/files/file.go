package file

import (
	"os"
	"path/filepath"

	"github.com/LiddleChild/findr/internal/core/engine"
	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/internal/models"
	"github.com/LiddleChild/findr/utils"
	"github.com/fatih/color"
)

type dirNode struct {
	path  string
	depth int
}

func Traverse(query string, arg *models.Argument) errorwrapper.ErrorWrapper {
	st := utils.NewStack[dirNode]()
	st.Push(dirNode{
		path:  arg.WorkingDirectory,
		depth: 0,
	})

	pattern := engine.CreatePattern(query)

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

			content := e.Name()
			if arg.ContentSearch && !e.IsDir() {
				bs, err := os.ReadFile(path)
				if err != nil {
					return errorwrapper.NewWithMessage(
						errorwrapper.Core,
						err,
						"error occured while reading files")
				}

				content = string(bs)
			}

			_, ok := pattern.Match(content)
			if ok {
				utils.HighlightedPrintln(content, query, color.FgRed)
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
