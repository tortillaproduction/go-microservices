package errs

type InternalError struct {
	message string
}

func (e *InternalError) Error() string {
	return e.message
}

func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}
