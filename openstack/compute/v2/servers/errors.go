package servers

import (
	"fmt"

	"github.com/rackspace/gophercloud"
)

type ErrNeitherImageIDNorImageNameProvided struct{}

func (e ErrNeitherImageIDNorImageNameProvided) Error() string {
	return "One and only one of the image ID and the image name must be provided."
}

type ErrNeitherFlavorIDNorFlavorNameProvided struct{}

func (e ErrNeitherFlavorIDNorFlavorNameProvided) Error() string {
	return "One and only one of the flavor ID and the flavor name must be provided."
}

type ErrInvalidHowParameterProvided struct{}

func (e ErrInvalidHowParameterProvided) Error() string {
	return "Unknown argument for 'how' parameter."
}

type ErrNoAdminPassProvided struct{}

func (e ErrNoAdminPassProvided) Error() string {
	return "You must provide an administrative password."
}

type ErrNoImageIDProvided struct{}

func (e ErrNoImageIDProvided) Error() string {
	return "You must provide an image ID."
}

type ErrNoIDProvided struct{}

func (e ErrNoIDProvided) Error() string {
	return "You must provide a server ID."
}

// ServerError is a generic error type for servers.
type ServerError struct {
	*gophercloud.UnexpectedResponseCodeError
	id string
}

func (se *ServerError) Error() string {
	return fmt.Sprintf("Error while executing HTTP request for server [%s]", se.id)
}

// Error404 overrides the generic 404 error message.
func (se *ServerError) Error404(e *gophercloud.UnexpectedResponseCodeError) error {
	se.UnexpectedResponseCodeError = e
	return &ServerNotFoundError{
		se,
	}
}

// ServerNotFoundError is an error type returned when a 404 is received during
// server HTTP operations.
type ServerNotFoundError struct {
	*ServerError
}

func (e *ServerNotFoundError) Error() string {
	return fmt.Sprintf("I couldn't find server [%s]", e.id)
}
