package errors

type DomainError interface {
	Code() string
	Message() string
	Error() string
}

type GenericError struct {
	code    string
	message string
}

func (e GenericError) Code() string    { return e.code }
func (e GenericError) Message() string { return e.message }
func (e GenericError) Error() string   { return e.code + ": " + e.message }
