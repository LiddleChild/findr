package options

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

case intensive search
directory / file things
concurrency
help

*/

type Metadata struct {
	Name        string
	Usage       string
	Description string
	Flags       []string
}

type Handler interface {
	Metadata() Metadata
	Handle(*core.Argument, []string) errorwrapper.ErrorWrapper
}
