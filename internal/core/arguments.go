package core

type Argument struct {
	MaxDepth         int
	ContentSearch    bool
	WorkingDirectory string
	IgnoredPaths     map[string]bool
}

func DefaultArgument() *Argument {
	return &Argument{
		MaxDepth:         5,
		ContentSearch:    false,
		WorkingDirectory: "",
		IgnoredPaths:     map[string]bool{},
	}
}
