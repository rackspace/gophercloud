package servers

import (
	"fmt"

	"github.com/rackspace/gophercloud"
)

const ErrNeitherImageIDNorImageNameProvided = "One and only one of the image ID and the image name must be provided."
const ErrNeitherFlavorIDNorFlavorNameProvided = "One and only one of the flavor ID and the flavor name must be provided."
const ErrInvalidHowParameterProvided = "Unknown argument for 'how' parameter"
const ErrNoAdminPassProvided = "You must provide an administrative password"
const ErrNoImageIDProvided = "You must provide an image ID"
const ErrNoIDProvided = "You must provide a server ID"

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

// ServerNotFoundError is an error type returned when a 404 is received during
// server HTTP operations.
type ServerNotFoundError struct {
	*gophercloud.UnexpectedResponseCodeError
	id string
}

func (e *ServerNotFoundError) Error() string {
	return fmt.Sprintf("I couldn't find server [%s]", e.id)
}
