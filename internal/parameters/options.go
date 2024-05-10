package parameters

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LiddleChild/findr/internal/arguments"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func MaxDepthOption(arg *arguments.Argument, args []string, cursor *int) errorwrapper.ErrorWrapper {
	i := *cursor
	val, err := strconv.Atoi(args[i+1])
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

	arg.SetMaxDepth(val)
	*cursor += 2

	return nil
}

func ContentSearchOption(arg *arguments.Argument, cursor *int) {
	arg.SetContentSearch(true)
	*cursor += 1
}

func WorkingDirectoryOption(arg *arguments.Argument, args []string, cursor *int) errorwrapper.ErrorWrapper {
	i := *cursor
	val := args[i+1]
	info, err := os.Stat(val)

	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Argument,
			err,
			fmt.Sprintf("-d: %v: no such directory", val))
	} else if !info.IsDir() {
		return errorwrapper.NewWithMessage(
			errorwrapper.Argument,
			err,
			fmt.Sprintf("-d: %v is not a directory", val))
	}

	arg.SetWorkingDirectory(val)
	*cursor += 2

	return nil
}
