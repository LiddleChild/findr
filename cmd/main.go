package main

import (
	"fmt"
	"os"

	"github.com/LiddleChild/findr/internal/cli"
	"github.com/LiddleChild/findr/internal/cli/options"
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: findr <query> [options]")
		os.Exit(0)
	}

	params := os.Args[1:]

	parser := cli.NewParser(
		&options.HelpOption{},
		&options.ContentSearchOption{},
		&options.IgnorePathOption{},
		&options.MaxDepthOption{},
		&options.WorkingDirectoryOption{},
		&options.CaseSensitiveOption{},
	)

	arg, werr := parser.Parse(params)
	if werr != nil {
		if werr.Type() != errorwrapper.Help {
			_, msg, _ := werr.Unwrap()
			fmt.Printf("%v\nuse `findr --help` for more informations\n", msg)
		}

		os.Exit(0)
	}

	werr = core.Traverse(arg)
	if werr != nil {
		_, msg, _ := werr.Unwrap()
		fmt.Println(msg)
		os.Exit(0)
	}
}
