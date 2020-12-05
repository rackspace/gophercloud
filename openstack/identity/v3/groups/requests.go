package groups

import (
	"github.com/rackspace/gophercloud"

	"github.com/rackspace/gophercloud/pagination"
)

// ListOpts allows you to query the List method.
type ListOpts struct {
	Name     string `q:"name"`
	DomainID string `q:"domain_id"`
}

// List enumerates the services available to a specific user.
func List(client *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	u := listURL(client)
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}
	u += q.String()
	createPage := func(r pagination.PageResult) pagination.Page {
		return GroupPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, u, createPage)
}
