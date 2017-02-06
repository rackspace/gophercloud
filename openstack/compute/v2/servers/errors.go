package servers

import (
	"fmt"
	"regexp"

	"github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"
)

type ErrNeitherImageIDNorImageNameProvided struct{ *gophercloud.InvalidInputError }

func (e ErrNeitherImageIDNorImageNameProvided) Error() string {
	return "One and only one of the image ID and the image name must be provided."
}

type ErrNeitherFlavorIDNorFlavorNameProvided struct{ *gophercloud.InvalidInputError }

func (e ErrNeitherFlavorIDNorFlavorNameProvided) Error() string {
	return "One and only one of the flavor ID and the flavor name must be provided."
}

type ErrInvalidHowParameterProvided struct{ *gophercloud.InvalidInputError }

func (e ErrInvalidHowParameterProvided) Error() string {
	return "Unknown argument for 'how' parameter."
}

type ErrNoAdminPassProvided struct{ *gophercloud.InvalidInputError }

func (e ErrNoAdminPassProvided) Error() string {
	return "You must provide an administrative password."
}

type ErrNoImageIDProvided struct{ *gophercloud.InvalidInputError }

func (e ErrNoImageIDProvided) Error() string {
	return "You must provide an image ID."
}

type ErrNoIDProvided struct{ *gophercloud.InvalidInputError }

func (e ErrNoIDProvided) Error() string {
	return "You must provide a server ID."
}

// ServerError is a generic error type for servers.
type ServerError struct {
	*gophercloud.UnexpectedResponseCodeError
	id string
}

func (se ServerError) Error() string {
	return fmt.Sprintf("Error while executing HTTP request for server [%s]", se.id)
}

type ErrInvalidImageIDProvided struct{ ServerError }

func (e ErrInvalidImageIDProvided) Error() string {
	return "Invalid imageRef provided"
}

// Error400 overrides the generic 400 error message.
func (se ServerError) Error400(e *gophercloud.UnexpectedResponseCodeError) error {
	se.UnexpectedResponseCodeError = e
	stringBody := string(e.Body)
	rxInvalidImageID := regexp.MustCompile(`Invalid imageRef provided`)
	switch {
	case rxInvalidImageID.MatchString(stringBody):
		return ErrInvalidImageIDProvided{se}
	}
	return se
}

type ErrPersonalityContentTooLong struct{ ServerError }

func (e ErrPersonalityContentTooLong) Error() string {
	return "Length of personality file content plus path exceeds 1000 bytes"
}

type ErrFlavorHasNoDisk struct{ ServerError }

func (e ErrFlavorHasNoDisk) Error() string {
	return "The flavor you selected doesn't have a disk on which to directly install an image"
}

// Error403 overrides the generic 403 error message.
func (se ServerError) Error403(e *gophercloud.UnexpectedResponseCodeError) error {
	se.UnexpectedResponseCodeError = e
	stringBody := string(e.Body)
	rxPersonalityContentTooLong := regexp.MustCompile(`Personality file content too long`)
	rxFlavorHasNoDisk := regexp.MustCompile(`compute_flavor:create:image_backed`)
	switch {
	case rxPersonalityContentTooLong.MatchString(stringBody):
		return ErrInvalidImageIDProvided{se}
	case rxFlavorHasNoDisk.MatchString(stringBody):
		return ErrFlavorHasNoDisk{se}
	}
	return se
}

// Error404 overrides the generic 404 error message.
func (se ServerError) Error404(e *gophercloud.UnexpectedResponseCodeError) error {
	se.UnexpectedResponseCodeError = e
	return &ServerNotFoundError{
		se,
	}
}

// ServerNotFoundError is an error type returned when a 404 is received during
// server HTTP operations.
type ServerNotFoundError struct {
	ServerError
}

func (e ServerNotFoundError) Error() string {
	return fmt.Sprintf("I couldn't find server [%s]", e.id)
}
