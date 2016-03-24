package groups

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToGroupListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	// UUID of the group at which you want to set a marker.
	Marker string `q:"marker"`

	// Integer value for the limit of values to return.
	Limit int `q:"limit"`
}

// ToGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToGroupListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)

	if err != nil {
		return "", err
	}

	return q.String(), nil
}

// List returns all scaling groups.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)

	if opts != nil {
		query, err := opts.ToGroupListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	createPageFn := func(r pagination.PageResult) pagination.Page {
		return GroupPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, url, createPageFn)
}
