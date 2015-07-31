package openstack

import (
	"fmt"

	tokens2 "github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

type ErrEndpointNotFound struct{}

func (e ErrEndpointNotFound) Error() string {
	return "No suitable endpoint could be found in the service catalog."
}

type ErrInvalidAvailabilityProvided struct{}

func (e ErrInvalidAvailabilityProvided) Error() string {
	return "Unexpected availability in endpoint query"
}

type ErrMultipleMatchingEndpointsV2 struct {
	endpoints []tokens2.Endpoint
}

func (e *ErrMultipleMatchingEndpointsV2) Error() string {
	return fmt.Sprintf("Discovered %d matching endpoints: %#v", len(e.endpoints), e.endpoints)
}

type ErrMultipleMatchingEndpointsV3 struct {
	endpoints []tokens3.Endpoint
}

func (e *ErrMultipleMatchingEndpointsV3) Error() string {
	return fmt.Sprintf("Discovered %d matching endpoints: %#v", len(e.endpoints), e.endpoints)
}
