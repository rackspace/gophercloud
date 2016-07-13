package extensions

import (
	"github.com/rackspace/gophercloud"
)

/*
AuthOptions stores information needed to authenticate to an OpenStack cluster.
Pass one to a provider's AuthenticatedClient function to authenticate and obtain a
ProviderClient representing an active session on that provider.

Its fields are the union of those recognized by each identity implementation and
provider.
*/
type AuthOptions struct {
	//Populate fields in gophercloud AuthOptions also.
	*gophercloud.AuthOptions

	// Trust allows users to authenticate with Trust ID,
	// The TrustID  field is to be used with Identity V3 API only.
	// ID of the Trust.
	TrustID string
}
