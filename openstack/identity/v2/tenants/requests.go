package tenants

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// ListOpts filters the Tenants that are returned by the List call.
type ListOpts struct {
	// Marker is the ID of the last Tenant on the previous page.
	Marker string `q:"marker"`

	// Limit specifies the page size.
	Limit int `q:"limit"`
}

// List enumerates the Tenants to which the current token has access.
func List(client *gophercloud.ServiceClient, opts *ListOpts) pagination.Pager {
	createPage := func(r pagination.PageResult) pagination.Page {
		return TenantPage{pagination.LinkedPageBase{PageResult: r}}
	}

	url := listURL(client)
	if opts != nil {
		q, err := gophercloud.BuildQueryString(opts)
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q.String()
	}

	return pagination.NewPager(client, url, createPage)
}

// EnabledState represents whether the tenant is enabled or not.
type EnabledState *bool

// Useful variables to use when creating or updating tenants.
var (
	iTrue  = true
	iFalse = false

	Enabled  EnabledState = &iTrue
	Disabled EnabledState = &iFalse
)

// CommonOpts are the parameters that are shared between CreateOpts and
// UpdateOpts
type CommonOpts struct {
	// A name is required. When provided, the value must be
	// unique or a 409 conflict error will be returned.
	Name string

	// Indicates whether this user is enabled or not.
	Enabled EnabledState

	// The description of this tenant.
	Description string
}

// CreateOpts represents the options needed when creating new tenants.
type CreateOpts CommonOpts

// CreateOptsBuilder describes struct types that can be accepted by the Create call.
type CreateOptsBuilder interface {
	ToTenantCreateMap() (map[string]interface{}, error)
}

// ToTenantCreateMap assembles a request body based on the contents of a CreateOpts.
func (opts CreateOpts) ToTenantCreateMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})

	if opts.Name != "" {
		m["name"] = opts.Name
	}
	if opts.Description != "" {
		m["description"] = opts.Description
	}
	if opts.Enabled != nil {
		m["enabled"] = &opts.Enabled
	}

	return map[string]interface{}{"tenant": m}, nil
}

// Create is the operation responsible for creating new tenants.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToTenantCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(rootURL(client), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return res
}

// Get requests details on a single tenant, either by ID.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(ResourceURL(client, id), &result.Body, nil)
	return result
}

// UpdateOptsBuilder allows extensions to add additional attributes to the Update request.
type UpdateOptsBuilder interface {
	ToTenantUpdateMap() map[string]interface{}
}

// UpdateOpts specifies the base attributes that may be updated on an existing tenant.
type UpdateOpts CommonOpts

// ToUserUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToTenantUpdateMap() map[string]interface{} {
	m := make(map[string]interface{})

	if opts.Name != "" {
		m["name"] = opts.Name
	}
	if opts.Description != "" {
		m["description"] = opts.Description
	}
	if opts.Enabled != nil {
		m["enabled"] = &opts.Enabled
	}

	return map[string]interface{}{"tenant": m}
}

// Update is the operation responsible for updating exist users by their UUID.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) UpdateResult {
	var result UpdateResult
	reqBody := opts.ToTenantUpdateMap()
	_, result.Err = client.Put(ResourceURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return result
}

// Delete is the operation responsible for permanently deleting an API user.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var result DeleteResult
	_, result.Err = client.Delete(ResourceURL(client, id), nil)
	return result
}
