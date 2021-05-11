package error

import (
	"errors"
	"net/http"

	pErrors "github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

// Error is used to pass an error during the request through the
// application with web specific context.
type Error struct {
	Err    error
	Status int
	Code   string
}

type response map[string]interface{}

func (err *Error) Error() string {
	return err.Err.Error()
}

//Wrap errors like pkg/errors
func Wrap(err error, msg string) {
	pErrors.Wrap(err, msg)
}

// NewRequestError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewRequestError(err interface{}, status int, code string) *Error {
	var e error
	switch v := err.(type) {
	case error:
		e = v
	case string:
		e = errors.New(v)
	}
	return &Error{Err: e, Status: status, Code: code}
}

// CustomHTTPErrorHandler for echo
var CustomHTTPErrorHandler = func(err error, c echo.Context) {
	// If the error was of the type *Error, the handler has
	// a specific status code and error to return.

	//todo keep structs for echo and custom error similar
	if webErr, ok := pErrors.Cause(err).(*Error); ok {
		response := response{"message": webErr.Err.Error(), "code": webErr.Code}
		c.JSON(webErr.Status, response)
	} else if echoError, ok := err.(*echo.HTTPError); ok {
		response := response{"message": echoError.Message, "code": echoError.Code}
		c.JSON(echoError.Code, response)
	} else {
		// If not, the handler sent any arbitrary error value so use 500.
		response := response{"message": http.StatusText(http.StatusInternalServerError), "code": "server_error"}
		c.JSON(500, response)
		c.Logger().Error(err)
	}
}
