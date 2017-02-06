package openstack

import (
	"fmt"

	"github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"
	tokens2 "github.com/rackspace/rack/internal/github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/rackspace/rack/internal/github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

type ErrEndpointNotFound struct{ *gophercloud.BaseError }

func (e ErrEndpointNotFound) Error() string {
	return "No suitable endpoint could be found in the service catalog."
}

type ErrInvalidAvailabilityProvided struct{ *gophercloud.InvalidInputError }

func (e ErrInvalidAvailabilityProvided) Error() string {
	return "Unexpected availability in endpoint query"
}

type ErrMultipleMatchingEndpointsV2 struct {
	*gophercloud.BaseError
	endpoints []tokens2.Endpoint
}

func (e *ErrMultipleMatchingEndpointsV2) Error() string {
	return fmt.Sprintf("Discovered %d matching endpoints: %#v", len(e.endpoints), e.endpoints)
}

type ErrMultipleMatchingEndpointsV3 struct {
	*gophercloud.BaseError
	endpoints []tokens3.Endpoint
}

func (e *ErrMultipleMatchingEndpointsV3) Error() string {
	return fmt.Sprintf("Discovered %d matching endpoints: %#v", len(e.endpoints), e.endpoints)
}

type ErrNoAuthURL struct{ *gophercloud.InvalidInputError }

func (e *ErrNoAuthURL) Error() string {
	return "Environment variable OS_AUTH_URL needs to be set."
}

type ErrNoUsername struct{ *gophercloud.InvalidInputError }

func (e *ErrNoUsername) Error() string {
	return "Environment variable OS_USERNAME needs to be set."
}

type ErrNoPassword struct{ *gophercloud.InvalidInputError }

func (e *ErrNoPassword) Error() string {
	return "Environment variable OS_PASSWORD needs to be set."
}
