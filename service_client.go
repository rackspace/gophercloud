package gophercloud

import (
	"strings"

	"github.com/racker/perigee"
)

// ServiceClient stores details required to interact with a specific service API implemented by a provider.
// Generally, you'll acquire these by calling the appropriate `New` method on a ProviderClient.
type ServiceClient struct {
	// Provider is a reference to the provider that implements this service.
	Provider *ProviderClient

	// Endpoint is the base URL of the service's API, acquired from a service catalog.
	// It MUST end with a /.
	Endpoint string
}

// ServiceURL constructs a URL for a resource belonging to this provider.
func (client *ServiceClient) ServiceURL(parts ...string) string {
	return client.Endpoint + strings.Join(parts, "/")
}

// RequestOptions contains extra information about how to perform a specific request.
type RequestOptions struct {
	ReqBody interface{}
	OkCodes []int
}

// Request performs an authenticated request against the service endpoint.
func (client *ServiceClient) Request(method, url string, options RequestOptions) CommonResult {
	result := CommonResult{}

	opts := perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		ReqBody:     options.ReqBody,
		OkCodes:     options.OkCodes,
		Results:     &result.Resp,
	}
	_, result.Err = perigee.Request(method, url, opts)
	return result
}
