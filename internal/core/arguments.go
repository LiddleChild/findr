package core

import "runtime"

type Argument struct {
	Query            string
	MaxDepth         int
	ContentSearch    bool
	WorkingDirectory string
	IgnoredPaths     map[string]bool
	CaseSensitive    bool
	NumWorkers       int
}

func DefaultArgument() *Argument {
	return &Argument{
		Query:            "",
		MaxDepth:         5,
		ContentSearch:    false,
		WorkingDirectory: "",
		IgnoredPaths:     map[string]bool{},
		CaseSensitive:    false,
		NumWorkers:       runtime.NumCPU(),
	}
}
