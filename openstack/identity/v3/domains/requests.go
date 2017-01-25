package domains

import (
	"errors"
	"github.com/rackspace/gophercloud"

	"github.com/rackspace/gophercloud/pagination"
)

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

// PairOpts allows you to pair roles, groups, on a domain
type PairOpts struct {
	ID      string `json:"id"`
	GroupID string `json:"group_id"`
	RoleID  string `json:"role_id"`
}

// Pair creates a relationship between a role, group, and domain
func Pair(client *gophercloud.ServiceClient, opts PairOpts) error {
	if opts.ID == "" || opts.GroupID == "" || opts.RoleID == "" {
		return errors.New("Domain, Role, and Group ids are required.")
	}

	reqOpts := &gophercloud.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Content-Type": ""},
	}

	var result PairResult
	_, result.Err = client.Put(pairDomainGroupAndRoleURL(client, opts.ID, opts.GroupID, opts.RoleID), nil, nil, reqOpts)
	return result.Err
}
