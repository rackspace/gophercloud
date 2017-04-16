package servers

import (
	"github.com/rackspace/gophercloud"
	os "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
)

// ListOpts allows pagination of the tenant's servers.
type ListOpts struct {
	Limit  int `q:"limit"`
	Offset int `q:"offset"`
}

// ToServerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List makes a request against the API to list servers accessible to you.
func List(client *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	return os.List(client, opts)
}
