package go_error_lib

const CodeFieldIsRequired = 1
const CodeFieldIsUnique = 2
const CodeFieldEntityNotFound = 3
const CodeFieldEntityHasExpired = 4
const CodeFieldIsInvalid = 5
const CodeFieldUnauthorized = 6
const CodeUnknown = 7

type Error interface {
	GetCode() int
	GetMessage() string
}

type ApplicationError struct {
	Err  error
	Code int
}

func NewApplicationError(err error, code int) ApplicationError {
	return ApplicationError{
		Err:  err,
		Code: code,
	}
}

func (a ApplicationError) GetCode() int {
	return a.Code
}

func (a ApplicationError) GetMessage() string {
	return a.Err.Error()
}

type DomainError struct {
	Err  error
	Code int
}

func NewDomainError(err error, code int) DomainError {
	return DomainError{
		Err:  err,
		Code: code,
	}
}

func (d DomainError) GetCode() int {
	return d.Code
}

func (d DomainError) GetMessage() string {
	return d.Err.Error()
}
