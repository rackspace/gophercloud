package gophercloud

import (
	"errors"
	"fmt"
)

var ErrTimeOut = errors.New("A time out occurred")

// BaseError is an error type that all other error types embed.
type BaseError struct {
	Type     error
	Function string
}

func (e *BaseError) Error() string {
	return e.Type.Error()
}

// InvalidInputError is an error type used for most non-HTTP Gophercloud errors.
type InvalidInputError struct {
	BaseError
	Argument string
	Value    interface{}
}

func (e *InvalidInputError) Error() string {
	return e.Type.Error()
}

// UnexpectedResponseCodeError is returned by the Request method when a response code other than
// those listed in OkCodes is encountered.
type UnexpectedResponseCodeError struct {
	BaseError
	URL      string
	Method   string
	Expected []int
	Actual   int
	Body     []byte
}

func (err *UnexpectedResponseCodeError) Error() string {
	return fmt.Sprintf(
		"Expected HTTP response code %v when accessing [%s %s], but got %d instead\n%s",
		err.Expected, err.Method, err.URL, err.Actual, err.Body,
	)
}

type defaultError401 struct {
	*UnexpectedResponseCodeError
}
type defaultError404 struct {
	*UnexpectedResponseCodeError
}
type defaultError405 struct {
	*UnexpectedResponseCodeError
}

func (e defaultError401) Error() string {
	return "Authentication failed"
}
func (e defaultError404) Error() string {
	return "Resource not found"
}
func (e defaultError405) Error() string {
	return "Method not allowed"
}

// Error401er is the interface resource error types implement to override the error message
// from a 401 error.
type Error401er interface {
	Error401(*UnexpectedResponseCodeError) error
}

// Error404er is the interface resource error types implement to override the error message
// from a 404 error.
type Error404er interface {
	Error404(*UnexpectedResponseCodeError) error
}

// Error405er is the interface resource error types implement to override the error message
// from a 405 error.
type Error405er interface {
	Error405(*UnexpectedResponseCodeError) error
}
