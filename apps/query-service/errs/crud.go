package errs

type CRUDError struct {
	message string
}

func (e *CRUDError) Error() string {
	return e.message
}

func NewCRUDError(message string) *CRUDError {
	return &CRUDError{message: message}
}
