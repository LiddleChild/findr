package options

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

type WorkingDirectoryOption struct{}

func (opt WorkingDirectoryOption) Metadata() Metadata {
	return Metadata{
		Name:        "working directory",
		Usage:       "(-d|--dir) <path>",
		Description: "set working directory",
		Flags:       []string{"-d", "--dir"},
	}
}

func (opt WorkingDirectoryOption) Handle(arg *core.Argument, values []string) errorwrapper.ErrorWrapper {
	if len(values) < 1 {
		return errorwrapper.New(
			errorwrapper.Parsing,
			errors.New("-d: require a path to directory"))
	}

	dir := values[0]
	info, err := os.Stat(dir)
	if err != nil {
		return errorwrapper.NewWithMessage(
			errorwrapper.Argument,
			err,
			fmt.Sprintf("-d: %v: no such directory", dir))
	} else if !info.IsDir() {
		return errorwrapper.NewWithMessage(
			errorwrapper.Argument,
			err,
			fmt.Sprintf("-d: %v is not a directory", dir))
	}

	arg.WorkingDirectory = path.Join(arg.WorkingDirectory, dir)

	return nil
}
