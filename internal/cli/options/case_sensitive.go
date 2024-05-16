package options

import (
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type CaseSensitiveOption struct{}

func (opt CaseSensitiveOption) Metadata() Metadata {
	return Metadata{
		Name:        "case sensitive",
		Usage:       "(-C|--case)",
		Description: "case sensitive query",
		Flags:       []string{"-C", "--case"},
	}
}

func (opt CaseSensitiveOption) Handle(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	arg.CaseSensitive = true
	return nil
}
