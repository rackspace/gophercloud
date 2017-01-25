package groups

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

type commonResult struct {
	gophercloud.Result
}

// Link the object to hold a project link.
type Link struct {
	Self string `json:"self,omitempty"`
}

// Group is main struct for holding group attributes.
type Group struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Links       Link   `json:"links"`
}

// GroupPage is a single page of Group results.
type GroupPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p GroupPage) IsEmpty() (bool, error) {
	groups, err := ExtractGroups(p)
	if err != nil {
		return true, err
	}
	return len(groups) == 0, nil
}

// ExtractGroups extracts a slice of Groups from a Collection acquired from List.
func ExtractGroups(page pagination.Page) ([]Group, error) {
	var response struct {
		Groups []Group `mapstructure:"groups"`
	}

	err := mapstructure.Decode(page.(GroupPage).Body, &response)
	return response.Groups, err
}
