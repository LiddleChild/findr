package options

import (
	"errors"
	"strconv"

	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func MaxDepthHandler(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
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
