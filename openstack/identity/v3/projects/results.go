package projects

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Project, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Project `json:"project"`
	}

	err := mapstructure.Decode(r.Body, &response)

	return &response.Project, err
}

// CreateResult Project create result
type CreateResult struct {
	commonResult
}

// GetResult Project details result
type GetResult struct {
	commonResult
}

// UpdateResult Project update result
type UpdateResult struct {
	commonResult
}

// DeleteResult Project delete result
type DeleteResult struct {
	commonResult
}

// Project Project struct
type Project struct {
	DomainID    string `mapstructure:"domain_id" json:"domain_id"`
	ParentID    string `mapstructure:"parent_id" json:"parent_id"`
	Enabled     bool   `mapstructure:"enabled" json:"enabled"`
	ID          string `mapstructure:"id" json:"id"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
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
