package options

import (
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type ContentSearchOption struct{}

func (opt ContentSearchOption) Metadata() Metadata {
	return Metadata{
		Name:        "content search",
		Usage:       "(-c|--content)",
		Description: "query file content instead of file name",
		Flags:       []string{"-c", "--content"},
	}
}

func (opt ContentSearchOption) Handle(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	arg.ContentSearch = true
	return nil
}
