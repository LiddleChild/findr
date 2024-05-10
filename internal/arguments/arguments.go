package arguments

type Argument struct {
	maxDepth      int
	contentSearch bool
	workingDir    string
}

func New() *Argument {
	return &Argument{
		maxDepth:      5,
		contentSearch: false,
		workingDir:    ".",
	}
}

func (cmd *Argument) SetMaxDepth(depth int) *Argument {
	cmd.maxDepth = depth
	return cmd
}

func (cmd *Argument) SetContentSearch(enb bool) *Argument {
	cmd.contentSearch = enb
	return cmd
}

func (cmd *Argument) SetWorkingDirectory(dir string) *Argument {
	cmd.workingDir = dir
	return cmd
}
