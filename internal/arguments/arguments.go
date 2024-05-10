package arguments

type Argument struct {
	MaxDepth         int
	ContentSearch    bool
	WorkingDirectory string
}

func New() *Argument {
	return &Argument{
		MaxDepth:         5,
		ContentSearch:    false,
		WorkingDirectory: ".",
	}
}
