package objects

import (
	"fmt"

	"github.com/rackspace/gophercloud"
)

// ObjectError is a generic error type for objects.
type ObjectError struct {
	*gophercloud.UnexpectedResponseCodeError
	container string
	name      string
}

func (oe *ObjectError) Error() string {
	return fmt.Sprintf("Error while executing HTTP request for object [%s] in container [%s]", oe.name, oe.container)
}

// Error404 overrides the generic 404 error message.
func (oe *ObjectError) Error404(e *gophercloud.UnexpectedResponseCodeError) error {
	oe.UnexpectedResponseCodeError = e
	return &ObjectNotFoundError{
		oe,
	}
}

// ObjectNotFoundError is an error type returned when a 404 is received during
// object HTTP operations.
type ObjectNotFoundError struct {
	*ObjectError
}

func (e *ObjectNotFoundError) Error() string {
	return fmt.Sprintf("I couldn't find object [%s] in container [%s]", e.name, e.container)
}
