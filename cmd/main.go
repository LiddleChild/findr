package main

import (
	"fmt"
	"os"

	"github.com/LiddleChild/findr/internal/cli"
	"github.com/LiddleChild/findr/internal/core"
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

	query, arg, werr := cli.Parse(params)
	if werr != nil {
		_, msg, _ := werr.Unwrap()
		fmt.Printf("%v\nuse `findr --help` for more informations\n", msg)
		os.Exit(0)
	}

	werr = core.Traverse(query, arg)
	if werr != nil {
		_, msg, _ := werr.Unwrap()
		fmt.Println(msg)
		os.Exit(0)
	}
}
