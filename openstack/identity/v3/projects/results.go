package projects

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// Project Project struct
type Project struct {
	// The ID of the domain for the project.
	DomainID string `mapstructure:"domain_id" json:"domain_id"`

	// The ID of the parent project.
	ParentID string `mapstructure:"parent_id" json:"parent_id"`

	// Enables or disables a project.
	// Set to true to enable the project or false to disable the project. Default is true.
	Enabled bool `mapstructure:"enabled" json:"enabled"`

	// The ID for the project.
	ID string `mapstructure:"id" json:"id"`

	// The project name.
	Name string `mapstructure:"name" json:"name"`

	// The project description.
	Description string `mapstructure:"description" json:"description"`
}

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Project, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Project Project `json:"project"`
	}

	err := mapstructure.Decode(r.Body, &response)

	return &response.Project, err
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of a update operation.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ProjectPage Page containing projects
type ProjectPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks if projects page is empty
func (p ProjectPage) IsEmpty() (bool, error) {
	projects, err := ExtractProjects(p)
	if err != nil {
		return true, err
	}
	return len(projects) == 0, nil
}

// ExtractProjects extracts projects list from response
func ExtractProjects(page pagination.Page) ([]Project, error) {
	var response struct {
		Projects []Project `mapstructure:"projects"`
	}

	err := mapstructure.Decode(page.(ProjectPage).Body, &response)

	return response.Projects, err
}
