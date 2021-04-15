package errors

// GenericHttpError is used to pass an error during the request through the application
type GenericHttpError struct {
	Status    int
	ErrorCode string
	ErrorMsg  error
}

// GenericHttpError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewGenericHttpError(status int, errorCode string, errorMsg error) error {
	return &GenericHttpError{status, errorCode, errorMsg}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (err *GenericHttpError) Error() string {
	return err.ErrorMsg.Error()
}
