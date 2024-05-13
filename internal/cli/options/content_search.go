package options

import (
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func ContentSearchHandler(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	arg.ContentSearch = true
	return nil
}
