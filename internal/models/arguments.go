package models

type Argument struct {
	MaxDepth         int
	ContentSearch    bool
	WorkingDirectory string
}

func DefaultArgument() *Argument {
	return &Argument{
		MaxDepth:         5,
		ContentSearch:    false,
		WorkingDirectory: ".",
	}
}
