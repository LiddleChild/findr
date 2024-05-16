package core

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/utils"
)

type dirNode struct {
	path  string
	depth int
}

var pattern *Pattern

func Traverse(arg *Argument) errorwrapper.ErrorWrapper {
	st := utils.NewStack[dirNode]()
	st.Push(dirNode{
		path:  arg.WorkingDirectory,
		depth: 0,
	})

	query := arg.Query
	if !arg.CaseSensitive {
		query = strings.ToLower(query)
	}
	pattern = CreatePattern(query)

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
			} else if !arg.ContentSearch {
				werr = HandleFilenameSearch(path, arg)
			}

			if werr != nil {
				return werr
			}
		}
	}

	return nil
}
