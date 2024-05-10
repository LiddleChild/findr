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
			val, err := strconv.Atoi(args[cursor+1])
			if err != nil {
				return "", nil, errorwrapper.NewWithMessage(
					errorwrapper.Argument,
					err,
					"-mx: value must be an integer")
			}
			
			if val < 0 {
				return "", nil, errorwrapper.NewWithMessage(
					errorwrapper.Argument,
					err,
					"-mx: value must be greater or equal to zero")
			}

			arg.SetMaxDepth(val)
			cursor += 2

		case "-c":
			arg.SetContentSearch(true)
			cursor += 1

		case "-d":
			val := args[cursor+1]
			info, err := os.Stat(val)

			if err != nil {
				return "", nil, errorwrapper.NewWithMessage(
					errorwrapper.Argument,
					err,
					fmt.Sprintf("-d: %v: no such directory", val))
			} else if !info.IsDir() {
				return "", nil, errorwrapper.NewWithMessage(
					errorwrapper.Argument,
					err,
					fmt.Sprintf("-d: %v is not a directory", val))
			}

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
