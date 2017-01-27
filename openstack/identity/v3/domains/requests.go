package domains

import (
	"github.com/rackspace/gophercloud"

	"github.com/rackspace/gophercloud/pagination"
)

type response struct {
	Domain Domain `json:"domains"`
}

// ListOpts allows you to query the List method.
type ListOpts struct {
	Name    string `q:"name"`
	Enabled string `q:"enabled"`
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
		return DomainPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, u, createPage)
}
