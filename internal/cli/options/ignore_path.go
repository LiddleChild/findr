package options

import (
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func IgnorePathHandler(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	ignores := make(map[string]bool)

	for _, path := range values {
		ignores[path] = true
	}

	arg.IgnoredPaths = ignores

	return nil
}
