package parameters

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LiddleChild/findr/internal/errorwrapper"
	"github.com/LiddleChild/findr/internal/models"
)

func MaxDepthOption(arg *models.Argument, args []string, cursor *int) errorwrapper.ErrorWrapper {
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

	arg.MaxDepth = val
	*cursor += 2

	return nil
}

func ContentSearchOption(arg *models.Argument, cursor *int) {
	arg.ContentSearch = true
	*cursor += 1
}

func WorkingDirectoryOption(arg *models.Argument, args []string, cursor *int) errorwrapper.ErrorWrapper {
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

	arg.WorkingDirectory = val
	*cursor += 2

	return nil
}
