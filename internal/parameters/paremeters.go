package parameters

import (
	"fmt"
	"os"
	"strings"

	"github.com/LiddleChild/findr/internal/arguments"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func Parse(args []string) (string, *arguments.Argument, errorwrapper.ErrorWrapper) {
	cursor := 0
	for cursor < len(args) && args[cursor][0] != '-' {
		cursor++
	}

	query := strings.Join(args[:cursor], " ")

	pwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	arg := arguments.New()
	arg.SetWorkingDirectory(pwd)

	for cursor < len(args) {
		switch args[cursor] {
		case "-mx":
			werr := MaxDepthOption(arg, args, &cursor)
			if werr != nil {
				return "", nil, werr
			}

		case "-c":
			ContentSearchOption(arg, &cursor)

		case "-d":
			WorkingDirectoryOption(arg, args, &cursor)

		default:
			return "", nil, errorwrapper.New(
				errorwrapper.Parameter,
				fmt.Errorf("unknown argument: %v", args[cursor]))
		}
	}

	return query, arg, nil
}
