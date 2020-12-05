package projects

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

// Project the object to hold a project.
type Project struct {
	ID          string `json:"id"`
	IsDomain    bool   `json:"is_domain"`
	Description string `json:"description"`
	DomainID    string `json:"domain_id"`
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	ParentID    string `json:"parent_id"`
	Links       Link   `json:"links"`
}

// Link the object to hold a project link.
type Link struct {
	Self string `json:"self"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a concrete Service.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Project, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Project `json:"project"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return &res.Project, err
}

// CreateResult the object to hold a project link.
type CreateResult struct {
	commonResult
}

// PairResult the object to error for failed pairs.
type PairResult struct {
	commonResult
}

// ProjectPage is a single page of Project results.
type ProjectPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p ProjectPage) IsEmpty() (bool, error) {
	projects, err := ExtractProjects(p)
	if err != nil {
		return true, err
	}
	return len(projects) == 0, nil
}

// ExtractProjects extracts a slice of Projects from a Collection acquired from List.
func ExtractProjects(page pagination.Page) ([]Project, error) {
	var response struct {
		Projects []Project `mapstructure:"projects"`
	}

	err := mapstructure.Decode(page.(ProjectPage).Body, &response)
	return response.Projects, err
}
