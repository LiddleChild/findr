package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type Parser struct {
	parsers map[string]OptionHandler
}

func NewParser(options ...OptionHandler) *Parser {
	parsers := map[string]OptionHandler{}

	for _, opt := range options {
		for _, name := range opt.Metadata().OptionNames {
			parsers[name] = opt
		}
	}

	return &Parser{parsers}
}

func (p Parser) Parse(args []string) (*core.Argument, errorwrapper.ErrorWrapper) {
	cursor := 0
	for cursor < len(args) && args[cursor][0] != '-' {
		cursor++
	}

	pwd, err := os.Getwd()
	if err != nil {
		return nil, errorwrapper.NewWithMessage(
			errorwrapper.Parsing,
			err,
			"error occured while getting current working directory")
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

		opt, ok := p.parsers[key]
		if !ok {
			return nil, errorwrapper.New(
				errorwrapper.Parsing,
				fmt.Errorf("unknown option: %v", key))
		}

		werr := opt.Handle(arg, value)
		if werr != nil {
			return nil, werr
		}
	}

	return arg, nil
}
