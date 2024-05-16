package options

import (
	"errors"
	"strconv"

	"github.com/LiddleChild/findr/internal/cli"
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type MaxDepthOption struct{}

func (opt MaxDepthOption) Metadata() cli.OptionMetadata {
	return cli.OptionMetadata{
		Name:        "max depth",
		Usage:       "(-mx|--max-depth) <depth>",
		Description: "max depth of traversing, 0 means working directory",
		OptionNames: []string{"-mx", "--max-depth"},
	}
}

func (opt MaxDepthOption) Handle(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	if len(values) < 1 {
		return errorwrapper.New(
			errorwrapper.Parsing,
			errors.New("-mx: require an integer value"))
	}

	val, err := strconv.Atoi(values[0])
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Argument,
			err,
			"-mx: value must be an integer")
	}

	if val < 0 {
		return errorwrapper.NewWithMessage(
			errorwrapper.Argument,
			err,
			"-mx: value must be greater or equal to zero")
	}

	arg.MaxDepth = val

	return nil
}
