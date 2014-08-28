package utils

import (
	"fmt"
	identity "github.com/rackspace/gophercloud/openstack/identity/v2"
)

// Client contains information that defines a generic Openstack Client.
type Client struct {
	// Endpoint is the URL against which to authenticate.
	Endpoint string
	// Authority holds the results of authenticating against the Endpoint.
	Authority identity.AuthResults
	// Options holds the authentication options. Useful for auto-reauthentication.
	Options identity.AuthOptions
}

// EndpointOpts contains options for finding an endpoint for an Openstack Client.
type EndpointOpts struct {
	// Type is the service type for the client (e.g., "compute", "object-store").
	// Type is a required field.
	Type string
	// Name is the service name for the client (e.g., "nova").
	// Name is not a required field, but it is used if present. Services can have the
	// same Type but different Name, which is one example of when both Type and Name are needed.
	Name string
	// Region is the region in which the service resides.
	Region string
	// URLType is they type of endpoint to be returned (e.g., "public", "private").
	// URLType is not required, and defaults to "public".
	URLType string
}

// NewClient returns a generic Openstack Client of type identity.Client. This is a helper function
// to create a client for the various Openstack services.
// Example (error checking omitted for brevity):
//		ao, err := utils.AuthOptions()
//		c, err := identity.NewClient(ao, identity.EndpointOpts{
//			Type: "compute",
//			Name: "nova",
//		})
//		serversClient := servers.NewClient(c.Endpoint, c.Authority, c.Options)
func NewClient(ao identity.AuthOptions, eo EndpointOpts) (Client, error) {
	client := Client{
		Options: ao,
	}

	ar, err := identity.Authenticate(ao)
	if err != nil {
		return client, err
	}

	client.Authority = ar

	sc, err := identity.GetServiceCatalog(ar)
	if err != nil {
		return client, err
	}

	ces, err := sc.CatalogEntries()
	if err != nil {
		return client, err
	}

	var eps []identity.Endpoint

	if eo.Name != "" {
		for _, ce := range ces {
			if ce.Type == eo.Type && ce.Name == eo.Name {
				eps = ce.Endpoints
			}
		}
	} else {
		for _, ce := range ces {
			if ce.Type == eo.Type {
				eps = ce.Endpoints
			}
		}
	}

	var rep string
	for _, ep := range eps {
		if ep.Region == eo.Region {
			switch eo.URLType {
			case "public":
				rep = ep.PublicURL
			case "private":
				rep = ep.InternalURL
			default:
				rep = ep.PublicURL
			}
		}
	}

	if rep != "" {
		client.Endpoint = rep
	} else {
		return client, fmt.Errorf("No endpoint for given service type (%s) name (%s) and region (%s)", eo.Type, eo.Name, eo.Region)
	}

	return client, nil
}