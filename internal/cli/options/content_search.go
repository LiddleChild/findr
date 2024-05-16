package options

import (
	"github.com/LiddleChild/findr/internal/cli"
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type ContentSearchOption struct{}

func (opt ContentSearchOption) Metadata() cli.OptionMetadata {
	return cli.OptionMetadata{
		Name:        "content search",
		Usage:       "(-c|--content)",
		Description: "query file content instead of file name",
		OptionNames: []string{"-c", "--content"},
	}
}

func (opt ContentSearchOption) Handle(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	arg.ContentSearch = true
	return nil
}
