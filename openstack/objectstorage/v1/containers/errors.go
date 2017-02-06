package containers

import (
	"fmt"

	"github.com/rackspace/gophercloud"
)

// ContainerError is a generic error type for containers.
type ContainerError struct {
	*gophercloud.UnexpectedResponseCodeError
	name string
}

func (ce *ContainerError) Error() string {
	return fmt.Sprintf("Error while executing HTTP request for container [%s]", ce.name)
}

// Error404 overrides the generic 404 error message.
func (ce *ContainerError) Error404(e *gophercloud.UnexpectedResponseCodeError) error {
	ce.UnexpectedResponseCodeError = e
	return &ContainerNotFoundError{
		ce,
	}
}

// ContainerNotFoundError is an error type returned when a 404 is received during
// container HTTP operations.
type ContainerNotFoundError struct {
	*ContainerError
}

func (e *ContainerNotFoundError) Error() string {
	return fmt.Sprintf("I couldn't find container [%s]", e.name)
}
