package gophercloud

import "fmt"

// BaseError is an error type that all other error types embed.
type BaseError struct {
	Function string
}

func (e *BaseError) Error() string {
	return "An error occurred while executing a Gophercloud request."
}

// InvalidInputError is an error type used for most non-HTTP Gophercloud errors.
type InvalidInputError struct {
	BaseError
	Argument string
	Value    interface{}
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("Invalid input provided for argument [%s]: [%+v]", e.Argument, e.Value)
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

type defaultError400 struct {
	*UnexpectedResponseCodeError
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
type defaultError408 struct {
	*UnexpectedResponseCodeError
}
type defaultError429 struct {
	*UnexpectedResponseCodeError
}
type defaultError500 struct {
	*UnexpectedResponseCodeError
}
type defaultError503 struct {
	*UnexpectedResponseCodeError
}

func (e defaultError400) Error() string {
	return "Invalid request due to incorrect syntax or missing required parameters."
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
func (e defaultError408) Error() string {
	return "The server timed out waiting for the request"
}
func (e defaultError429) Error() string {
	return "Too many requests have been sent in a given amount of time. Pause requests, wait up to one minute, and try again."
}
func (e defaultError500) Error() string {
	return "Internal Server Error"
}
func (e defaultError503) Error() string {
	return "The service is currently unable to handle the request due to a temporary overloading or maintenance. This is a temporary condition. Try again later."
}

// Error400er is the interface resource error types implement to override the error message
// from a 400 error.
type Error400er interface {
	Error400(*UnexpectedResponseCodeError) error
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

// Error408er is the interface resource error types implement to override the error message
// from a 408 error.
type Error408er interface {
	Error408(*UnexpectedResponseCodeError) error
}

// Error429er is the interface resource error types implement to override the error message
// from a 429 error.
type Error429er interface {
	Error429(*UnexpectedResponseCodeError) error
}

// Error500er is the interface resource error types implement to override the error message
// from a 500 error.
type Error500er interface {
	Error500(*UnexpectedResponseCodeError) error
}

// Error503er is the interface resource error types implement to override the error message
// from a 503 error.
type Error503er interface {
	Error503(*UnexpectedResponseCodeError) error
}

type ErrTimeOut struct{ *BaseError }

func (e *ErrTimeOut) Error() string {
	return "A time out occurred"
}

type ErrUnableToReauthenticate struct {
	*BaseError
	OriginalError error
}

func (e *ErrUnableToReauthenticate) Error() string {
	return fmt.Sprintf("Unable to re-authenticate: %s", e.OriginalError)
}

type ErrErrorAfterReauthentication struct {
	*BaseError
	OriginalError error
}

func (e *ErrErrorAfterReauthentication) Error() string {
	return fmt.Sprintf("Successfully re-authenticated, but got error executing request: %s", e.OriginalError)
}
