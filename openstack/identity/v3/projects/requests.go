package projects

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type projectOpts struct {
	DomainID    string
	ParentID    string
	Enabled     *bool
	Name        string
	Description string
}

// CreateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Create operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type CreateOptsBuilder interface {
	ToProjectCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts projectOpts

// ToProjectCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToProjectCreateMap() (map[string]interface{}, error) {
	p := make(map[string]interface{})

	if opts.DomainID != "" {
		p["domain_id"] = opts.DomainID
	}
	if opts.ParentID != "" {
		p["parent_id"] = opts.ParentID
	}
	if opts.Enabled != nil {
		p["enabled"] = &opts.Enabled
	}
	if opts.Name != "" {
		p["name"] = opts.Name
	}
	if opts.Description != "" {
		p["description"] = opts.Description
	}

	return map[string]interface{}{"project": p}, nil
}

// Create accepts a CreateOpts struct and creates a new project using the values
// provided.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// network. An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToProjectCreateMap()
	if err != nil {
		res.Err = err
		return res
	}
	_, res.Err = client.Post(createURL(client), reqBody, &res.Body, nil)
	return res
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

// ListOpts allows the filtering and of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to
// the project attributes you want to filter on. Page and PerPage are used for
// pagination.
type ListOpts struct {
	DomainID string `q:"domain_id"`
	ParentID string `q:"parent_id"`
	Name     string `q:"name"`
	Enabled  bool   `q:"enabled"`
	Page     int    `q:"page"`
	PerPage  int    `q:"per_page"`
}

// ToProjectListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProjectListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// projects. It accepts a ListOpts struct, which allows you to filter the
// returned collection.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToProjectListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ProjectPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific project based on its unique ID.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = client.Get(getURL(client, id), &res.Body, nil)
	return res
}

// UpdateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Update operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type UpdateOptsBuilder interface {
	ToProjectUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts projectOpts

// ToProjectUpdateMap casts a UpdateOpts struct to a map.
func (opts UpdateOpts) ToProjectUpdateMap() (map[string]interface{}, error) {
	p := make(map[string]interface{})

	if opts.DomainID != "" {
		p["domain_id"] = opts.DomainID
	}
	if opts.ParentID != "" {
		p["parent_id"] = opts.ParentID
	}
	if opts.Enabled != nil {
		p["enabled"] = &opts.Enabled
	}
	if opts.Name != "" {
		p["name"] = opts.Name
	}
	if opts.Description != "" {
		p["description"] = opts.Description
	}

	return map[string]interface{}{"project": p}, nil
}

// Update accepts a UpdateOpts struct and updates an existing project using the
// values provided. For more information, see the Create function.
func Update(client *gophercloud.ServiceClient, projectID string, opts UpdateOptsBuilder) UpdateResult {
	var res UpdateResult

	reqBody, err := opts.ToProjectUpdateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Patch(updateURL(client, projectID), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

// Delete accepts a unique ID and deletes the project associated with it.
func Delete(client *gophercloud.ServiceClient, projectID string) DeleteResult {
	var result DeleteResult
	_, result.Err = client.Delete(deleteURL(client, projectID), nil)
	return result
}
