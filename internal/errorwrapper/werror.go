package errorwrapper

type WrapperType string

const (
	Argument  WrapperType = "Argument"
	Parameter WrapperType = "Parameter"
)

type ErrorWrapper interface {
	Unwrap() (WrapperType, string, error)
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