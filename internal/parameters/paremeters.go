package parameters

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LiddleChild/findr/internal/arguments"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func Parse(args []string) (string, *arguments.Argument, errorwrapper.ErrorWrapper) {
	cursor := 0
	for args[cursor][0] != '-' {
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
			val, err := strconv.Atoi(args[cursor + 1])
			if err != nil {
				return "", nil, errorwrapper.NewWithMessage(
					errorwrapper.Argument,
					err,
					"-mx: value must be an integer")
			}

			arg.SetMaxDepth(val)
			cursor += 2

		case "-c":
			arg.SetContentSearch(true)
			cursor += 1

		case "-d":
			val := args[cursor + 1]
			arg.SetWorkingDirectory(val)
			cursor += 2

		default:
			return "", nil, errorwrapper.New(
				errorwrapper.Parameter,
				fmt.Errorf("unknown argument: %v", args[cursor]))
		}
	}
	
	return query, arg, nil
}