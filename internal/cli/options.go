package cli

import (
	"github.com/LiddleChild/findr/internal/cli/options"
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

type OptionHandler func(*core.Argument, []string) errorwrapper.ErrorWrapper

var MappedOptionHandler = map[string]OptionHandler{
	"-mx": options.MaxDepthHandler,
	"-c":  options.ContentSearchHandler,
	"-d":  options.WorkingDirectoryHandler,
}
