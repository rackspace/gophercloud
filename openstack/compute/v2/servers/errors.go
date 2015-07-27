package servers

import (
	"fmt"

	"github.com/rackspace/gophercloud"
)

// ServerNotFoundError is an error type returned when a 404 is received during
// server HTTP operations.
type ServerNotFoundError struct {
	*gophercloud.UnexpectedResponseCodeError
	id string
}

func (snfe *ServerNotFoundError) Error() string {
	return fmt.Sprintf("I couldn't find server [%s]", snfe.id)
}

// ServerError is a generic error type for servers.
type ServerError struct {
	id string
}

func (se *ServerError) Error() string {
	return fmt.Sprintf("Error while executing HTTP request for server [%s]", se.id)
}

// Error404 overrides the generic 404 error message.
func (se *ServerError) Error404(e *gophercloud.UnexpectedResponseCodeError) error {
	return &ServerNotFoundError{
		e,
		se.id,
	}
}
