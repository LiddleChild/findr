package main

import (
	"fmt"
	"os"

	"github.com/LiddleChild/findr/internal/parameters"
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
}