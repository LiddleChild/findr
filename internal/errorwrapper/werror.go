package errorwrapper

type WrapperType string

const (
	Parsing  WrapperType = "Parsing"
	Help     WrapperType = "Help"
	Argument WrapperType = "Argument"
	Core     WrapperType = "Core"
)

type ErrorWrapper interface {
	Unwrap() (WrapperType, string, error)
	Type() WrapperType
	Message() string
	Error() error
}

type errorWrapper struct {
	errorType WrapperType
	message   string
	error     error
}

func New(errorType WrapperType, error error) ErrorWrapper {
	return NewWithMessage(errorType, error, error.Error())
}

func NewWithMessage(errorType WrapperType, error error, message string) ErrorWrapper {
	return &errorWrapper{
		errorType,
		message,
		error,
	}
}

func (err *errorWrapper) Unwrap() (WrapperType, string, error) {
	return err.errorType, err.message, err.error
}

func (err *errorWrapper) Type() WrapperType {
	return err.errorType
}

func (err *errorWrapper) Message() string {
	return err.message
}

func (err *errorWrapper) Error() error {
	return err.error
}
