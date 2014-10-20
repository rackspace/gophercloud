package client

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/testhelper"
)

// Fake token to use.
const TokenID = "cbc36478b0bd8e67e89469c7749d4127"

// ServiceClient returns a generic service client for use in tests.
func ServiceClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		Provider: &gophercloud.ProviderClient{TokenID: TokenID},
		Endpoint: testhelper.Endpoint(),
	}
}