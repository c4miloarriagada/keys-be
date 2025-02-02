package errors

type ValidationError struct {
	GenericError
}

func NewValidationError(code string, message string) *ValidationError {
	return &ValidationError{
		GenericError: GenericError{
			code:    code,
			message: message,
		},
	}
}

type NotFoundError struct {
	GenericError
}

func NewNotFoundError(code string, message string) *NotFoundError {
	return &NotFoundError{
		GenericError: GenericError{
			code:    code,
			message: message,
		},
	}
}

type UnauthenticatedError struct {
	GenericError
}

func NewUnauthenticatedError(code string, message string) *UnauthenticatedError {
	return &UnauthenticatedError{
		GenericError: GenericError{
			code:    code,
			message: message,
		},
	}
}

type InternalServerError struct {
	GenericError
}

func NewInternalServerError(code string, message string) *InternalServerError {
	return &InternalServerError{
		GenericError: GenericError{
			code:    code,
			message: message,
		},
	}
}
