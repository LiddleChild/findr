package core

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/utils"
)

type dirNode struct {
	path  string
	depth int
}

type pathData struct {
	path  string
	isDir bool
}

var pattern *Pattern

func Traverse(arg *Argument) errorwrapper.ErrorWrapper {
	query := arg.Query
	if !arg.CaseSensitive {
		query = strings.ToLower(query)
	}

	pattern = CreatePattern(query)

	done := make(chan struct{})
	werrCh := make(chan errorwrapper.ErrorWrapper, 1)
	pathCh := walkerRoutine(arg, done, werrCh)

	var wg sync.WaitGroup

	for range arg.NumWorkers {
		wg.Add(1)
		go func() {
			workerRoutine(arg, pathCh, done, werrCh)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case <-done:
			return nil
		case werr := <-werrCh:
			if werr != nil {
				return werr
			}
		}
	}
}

func workerRoutine(arg *Argument, pathCh <-chan pathData, done chan struct{}, werrCh chan<- errorwrapper.ErrorWrapper) {
	for path := range pathCh {
		select {
		case <-done:
			return
		default:
			var werr errorwrapper.ErrorWrapper
			if arg.ContentSearch && !path.isDir {
				werr = handleContentSearch(path.path, arg)
			} else if !arg.ContentSearch {
				werr = HandleFilenameSearch(path.path, arg)
			}

			werrCh <- werr
		}
	}
}

func walkerRoutine(arg *Argument, done <-chan struct{}, werrCh chan<- errorwrapper.ErrorWrapper) <-chan pathData {
	pathCh := make(chan pathData)

	go func() {
		defer close(pathCh)
		werrCh <- walk(arg, func(path string, isDir bool) errorwrapper.ErrorWrapper {
			select {
			case pathCh <- pathData{path, isDir}:
			case <-done:
				return errorwrapper.New(errorwrapper.Core, errors.New("done"))
			}

			return nil
		})
	}()

	return pathCh
}

func walk(arg *Argument, handler func(path string, isDir bool) errorwrapper.ErrorWrapper) errorwrapper.ErrorWrapper {
	st := utils.NewStack[dirNode]()
	st.Push(dirNode{
		path:  arg.WorkingDirectory,
		depth: 0,
	})

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

			werr := handler(path, e.IsDir())
			if werr != nil {
				return werr
			}
		}
	}

	return nil
}
