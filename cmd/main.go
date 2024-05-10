package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/LiddleChild/findr/internal/parameters"
	"github.com/LiddleChild/findr/utils"
)

/*

usage: findr <query> <options>

options
-help
-mx <max-depth>  : set max directory depth defaults to 5
-c               : search for content in files
-d <dir>         : set search root directory


*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: findr <query> <options>")
		os.Exit(0)
	}

	params := os.Args[1:]

	query, args, werr := parameters.Parse(params)
	if werr != nil {
		_, msg, _ := werr.Unwrap()
		fmt.Printf("%v\nuse `findr --help` for more informations\n", msg)
		os.Exit(0)
	}

	fmt.Println(query, args)

	type FileStack struct {
		path  string
		depth int
	}

	st := utils.NewStack[FileStack]()
	st.Push(FileStack{
		path: args.WorkingDirectory,
		depth: 0,
	})

	for st.Size() > 0 {
		top := st.Top()
		st.Pop()

		if top.depth > args.MaxDepth {
			continue
		}

		entries, err := os.ReadDir(top.path)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		for _, e := range entries {
			if strings.Contains(e.Name(), query) {
				fmt.Println(filepath.Join(top.path, e.Name()))
			}

			if e.IsDir() {
				st.Push(FileStack{
					path: filepath.Join(top.path, e.Name()),
					depth: top.depth + 1,
				})
			}
		}
	}
}
