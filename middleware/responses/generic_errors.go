package responses

import (
	"errors"
	"net/http"
)

// GenericBadRequestError implements a generic not found error. It receives the code and message as params.
func GenericBadRequestError(code string, message string) error {
	return NewGenericHttpError(http.StatusBadRequest, code, errors.New(message))
}

// GenericAlreadyExistsError implements a generic not found error. It receives the code and message as params.
func GenericAlreadyExistsError(code string, message string) error {
	return NewGenericHttpError(http.StatusBadRequest, code, errors.New(message))
}

// GenericNotFoundError implements a generic not found error. It receives the code and message as params.
func GenericNotFoundError(code string, message string) error {
	return NewGenericHttpError(http.StatusNotFound, code, errors.New(message))
}

// GenericInternalServerError implements a generic not found error. It receives the code and message as params.
func GenericInternalServerError(code string, message string) error {
	return NewGenericHttpError(http.StatusInternalServerError, code, errors.New(message))
}
