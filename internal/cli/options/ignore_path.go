package options

import (
	"github.com/LiddleChild/findr/internal/cli"
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type IgnorePathOption struct{}

func (opt IgnorePathOption) Metadata() cli.OptionMetadata {
	return cli.OptionMetadata{
		Name:        "ignore path",
		Usage:       "(-i|--ignore) [path(s)]",
		Description: "specified path(s) will be ignored while traversing, whitespace is used to separate between multiple paths",
		OptionNames: []string{"-i", "--ignore"},
	}
}

func (opt IgnorePathOption) Handle(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	ignores := make(map[string]bool)

	for _, path := range values {
		ignores[path] = true
	}

	arg.IgnoredPaths = ignores

	return nil
}
