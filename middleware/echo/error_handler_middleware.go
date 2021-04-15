package middleware

import (
	"github.com/labstack/echo/v4"
	e "github.com/laironacosta/kit-go/middleware"
	"github.com/pkg/errors"
	"net/http"
)

type (
	ErrorHandlerMiddlewareInterface interface {
		HandlerError(next echo.HandlerFunc) echo.HandlerFunc
	}

	ErrorHandlerMiddleware struct {
	}

	// Map defines a generic map of type `map[string]interface{}`.
	Map map[string]interface{}
)

func NewErrorHandlerMiddleware() ErrorHandlerMiddlewareInterface {
	return &ErrorHandlerMiddleware{}
}

func (h *ErrorHandlerMiddleware) HandlerError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			var er error
			var erCode string
			var status int
			// Build the error response
			switch act := errors.Cause(err).(type) {
			case *e.GenericHttpError:
				er = act.ErrorMsg
				erCode = act.ErrorCode
				status = act.Status
			default:
				er = errors.New(http.StatusText(http.StatusInternalServerError))
				erCode = "server_error"
				status = http.StatusInternalServerError
			}
			return c.JSON(status, Map{"error_code": erCode, "error_message": er.Error()})
		}
		return nil
	}
}
