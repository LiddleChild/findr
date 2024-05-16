package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func Parse(args []string) (*core.Argument, errorwrapper.ErrorWrapper) {
	cursor := 0
	for cursor < len(args) && args[cursor][0] != '-' {
		cursor++
	}

	pwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	arg := core.DefaultArgument()
	arg.WorkingDirectory = pwd

	arg.Query = strings.Join(args[:cursor], " ")
	if len(arg.Query) == 0 {
		return nil, errorwrapper.New(
			errorwrapper.Parsing,
			errors.New("query cannot be empty"))
	}

	for cursor < len(args) {
		key := args[cursor]

		cursor++
		start := cursor
		for cursor < len(args) && args[cursor][0] != '-' {
			cursor++
		}

		value := args[start:cursor]

		handler, ok := MappedOptionHandler[key]
		if !ok {
			return nil, errorwrapper.New(
				errorwrapper.Parsing,
				fmt.Errorf("unknown option: %v", key))
		}

		werr := handler(arg, value)
		if werr != nil {
			return nil, werr
		}
	}

	return arg, nil
}
