package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorCode holds the type of errors
type ErrorCode string

const (
	InternalError    ErrorCode = "Internal Error"
	NotFound         ErrorCode = "Not Found"
	InvalidArgument  ErrorCode = "Invalid Argument"
	Unauthenticated  ErrorCode = "Unauthenticated"
	PermissionDenied ErrorCode = "Permission Denied"
	Unknown          ErrorCode = "Unknown"
)

type customError struct {
	code           ErrorCode
	originalError  error
	displayMessage string
}

// New creates a new customError
func (code ErrorCode) New(msg string) error {
	return customError{code: code, originalError: errors.New(msg)}
}

// Newf creates a new customError with formatted message
func (code ErrorCode) Newf(msg string, args ...interface{}) error {
	return customError{code: code, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error
func (code ErrorCode) Wrap(err error, msg string) error {
	return code.Wrapf(err, msg)
}

// Wrapf creates a new wrapped error with formatted message
func (code ErrorCode) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{code: code, originalError: errors.Wrapf(err, msg, args...)}
}

// New creates a no type error
func New(msg string) error {
	return customError{code: InternalError, originalError: errors.New(msg)}
}

// Newf creates an InternalError error with formatted message
func Newf(msg string, args ...interface{}) error {
	return customError{
		code:          InternalError,
		originalError: errors.New(fmt.Sprintf(msg, args...)),
	}
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			code:          customErr.code,
			originalError: wrappedError,
		}
	}

	return customError{code: InternalError, originalError: wrappedError}
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// WithDisplayMessage returns a error containing a display message
func WithDisplayMessage(err error, msg string) error {
	if customErr, ok := err.(customError); ok {
		return customError{
			code:           customErr.code,
			originalError:  err,
			displayMessage: msg,
		}
	}

	return customError{code: InternalError, originalError: err, displayMessage: msg}
}

// Code retrives the error code from an error, defaults to InternalError
func Code(err error) ErrorCode {
	if customErr, ok := err.(customError); ok {
		return customErr.code
	}

	return InternalError
}

// Cause retrives the original error
// Note that it will return the error created internally from github.com/pkg/errors
func Cause(err error) error {
	custom, ok := err.(customError)
	if ok {
		return Cause(errors.Cause(custom.originalError))
	}

	return errors.Cause(err)
}

// DisplayMessage retrives the display message
func DisplayMessage(err error) string {
	custom, ok := err.(customError)
	if ok {
		if custom.displayMessage != "" {
			return custom.displayMessage
		}
		return string(custom.code)
	}

	return string(InternalError)
}

// Error returns the mssage of a customError
func (error customError) Error() string {
	return error.originalError.Error()
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// StackTrace retrives the stack trace of an error
func StackTrace(err error) errors.StackTrace {
	custom, ok := err.(customError)
	if ok {
		return StackTrace(custom.originalError)
	}

	return err.(stackTracer).StackTrace()[1:]
}
