package projects

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type project struct {
	DomainID    string `json:"domain_id"`
	ParentID    string `json:"parent_id"`
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ProjectOpts Options for project create & update
type ProjectOpts struct {
	DomainID    string
	ParentID    string
	Name        string
	Enabled     bool
	Description string
}

// Create Creates project
func Create(client *gophercloud.ServiceClient, opts ProjectOpts) CreateResult {
	type request struct {
		Project project `json:"project"`
	}

	reqBody := request{
		Project: project{
			DomainID:    opts.DomainID,
			ParentID:    opts.ParentID,
			Name:        opts.Name,
			Enabled:     opts.Enabled,
			Description: opts.Description,
		},
	}
	var result CreateResult
	_, result.Err = client.Post(listURL(client), reqBody, &result.Body, nil)
	return result
}

// ListOpts Options for listing projects
type ListOpts struct {
	DomainID string `q:"domain_id"`
	ParentID string `q:"parent_id"`
	Name     string `q:"name"`
	Enabled  bool   `q:"enabled"`
	Page     int    `q:"page"`
	PerPage  int    `q:"per_page"`
}

// List Lists projects
func List(client *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(client)
	query, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	url += query.String()
	createPage := func(r pagination.PageResult) pagination.Page {
		return ProjectPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, url, createPage)
}

// Get Shows project details
func Get(client *gophercloud.ServiceClient, projectID string) GetResult {
	var result GetResult
	_, result.Err = client.Get(projectURL(client, projectID), &result.Body, nil)
	return result
}

// Update Updates project information
func Update(client *gophercloud.ServiceClient, projectID string, opts ProjectOpts) UpdateResult {
	type request struct {
		Project project `json:"project"`
	}

	reqBody := request{
		Project: project{
			DomainID:    opts.DomainID,
			ParentID:    opts.ParentID,
			Name:        opts.Name,
			Enabled:     opts.Enabled,
			Description: opts.Description,
		},
	}
	var result UpdateResult
	_, result.Err = client.Put(projectURL(client, projectID), reqBody, &result.Body, nil)
	return result
}

// Delete Deletes project
func Delete(client *gophercloud.ServiceClient, projectID string) DeleteResult {
	var result DeleteResult
	_, result.Err = client.Delete(projectURL(client, projectID), nil)
	return result
}
