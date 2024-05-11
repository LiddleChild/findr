package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/LiddleChild/findr/internal/core/engine"
	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/internal/models"
	"github.com/LiddleChild/findr/utils"
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
			if _, ok := pattern.Match(e.Name()); ok {
				fmt.Println(filepath.Join(dir.path, e.Name()))
			}

			if e.IsDir() && dir.depth < arg.MaxDepth {
				st.Push(dirNode{
					path:  filepath.Join(dir.path, e.Name()),
					depth: dir.depth + 1,
				})
			}
		}
	}

	return nil
}
