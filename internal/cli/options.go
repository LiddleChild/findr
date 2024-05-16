package cli

import (
	"github.com/LiddleChild/findr/internal/core"
	"github.com/LiddleChild/findr/internal/errorwrapper"
)

/*

usage: findr <query> <options>

options
-help
-mx <max-depth>  : set max directory depth defaults to 5
-c               : search for content in files
-d <dir>         : set search root directory


*/

type OptionMetadata struct {
	Name        string
	Usage       string
	Description string
	OptionNames []string
}

type OptionHandler interface {
	Metadata() OptionMetadata
	Handle(*core.Argument, []string) errorwrapper.ErrorWrapper
}
