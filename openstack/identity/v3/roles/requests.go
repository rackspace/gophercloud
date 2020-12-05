package roles

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// ListAssignmentsOptsBuilder allows extensions to add additional parameters to
// the ListAssignments request.
type ListAssignmentsOptsBuilder interface {
	ToRolesListAssignmentsQuery() (string, error)
}

// ListAssignmentsOpts allows you to query the ListAssignments method.
// Specify one of or a combination of GroupId, RoleId, ScopeDomainId, ScopeProjectId,
// and/or UserId to search for roles assigned to corresponding entities.
// Effective lists effective assignments at the user, project, and domain level,
// allowing for the effects of group membership.
type ListAssignmentsOpts struct {
	GroupId        string `q:"group.id"`
	RoleId         string `q:"role.id"`
	ScopeDomainId  string `q:"scope.domain.id"`
	ScopeProjectId string `q:"scope.project.id"`
	UserId         string `q:"user.id"`
	Effective      bool   `q:"effective"`
}

// ToRolesListAssignmentsQuery formats a ListAssignmentsOpts into a query string.
func (opts ListAssignmentsOpts) ToRolesListAssignmentsQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// ListAssignments enumerates the roles assigned to a specified resource.
func ListAssignments(client *gophercloud.ServiceClient, opts ListAssignmentsOptsBuilder) pagination.Pager {
	url := listAssignmentsURL(client)
	query, err := opts.ToRolesListAssignmentsQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}
	url += query
	createPage := func(r pagination.PageResult) pagination.Page {
		return RoleAssignmentsPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, url, createPage)
}

// CreateOpts allows you to create a role
type CreateOpts struct {
	DomainID string `json:"domain_id"`
	Name     string `json:"name"`
}

// Create adds a new role using the provieded client.
func Create(client *gophercloud.ServiceClient, opts CreateOpts) CreateResult {
	type request struct {
		Role CreateOpts `json:"role"`
	}

	req := request{Role: opts}

	var result CreateResult
	_, result.Err = client.Post(listURL(client), req, &result.Body, nil)
	return result
}

// ListOpts allows you to query the List method.
type ListOpts struct {
	Name     string `q:"name"`
	DomainID string `q:"domain_id"`
}

// List enumerates the roles available to a specific user.
func List(client *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	u := listURL(client)
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}
	u += q.String()
	createPage := func(r pagination.PageResult) pagination.Page {
		return RolePage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, u, createPage)
}

