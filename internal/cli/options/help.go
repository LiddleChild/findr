package options

import (
	"errors"
	"fmt"
	"strings"

	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type HelpOption struct {
	Handlers []Handler
}

func (opt HelpOption) Metadata() Metadata {
	return Metadata{
		Name:        "help",
		Usage:       "--help",
		Description: "show option descriptions",
		Flags:       []string{"--help"},
	}
}

func (opt HelpOption) Handle(_ *core.Argument, _ []string) errorwrapper.ErrorWrapper {
	fmt.Println("Available options")

	for _, h := range opt.Handlers {
		fmt.Printf("%-24s`%s`: %s\n%29s: %s\n\n",
			strings.Join(h.Metadata().Flags, " | "),
			h.Metadata().Name,
			h.Metadata().Description,
			"usage",
			h.Metadata().Usage)
	}

	return errorwrapper.New(
		errorwrapper.Help,
		errors.New("help"))
}
