package groups

import (
	"github.com/mitchellh/mapstructure"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type groupResult struct {
	gophercloud.Result
}

// Group represents an Auto Scale group.
type Group struct {
	// UUID for the group
	ID string `mapstructure:"id" json:"id"`

	// State information for the group
	State State `mapstructure:"state" json:"state"`
}

// State represents the state information belonging to an Auto Scale group.
// TODO: Document these things.
type State struct {
	Name string `mapstructure:"name" json:"name"`

	Status Status `mapstructure:"status" json:"status"`

	DesiredCapacity int `mapstructure:"desiredCapacity" json:"desiredCapacity"`

	PendingCapacity int `mapstructure:"pendingCapacity" json:"pendingCapacity"`

	ActiveCapacity int `mapstructure:"activeCapacity" json:"activeCapacity"`

	Paused bool `mapstructure:"paused" json:"paused"`

	Active []map[string]interface{} `mapstructure:"active" json:"active"`
}

// Status indicates the status of an Auto Scale group.
// TODO: Find out more about this.
type Status string

const (
	// ACTIVE indicates something, but I'm not sure what because it's not
	// documented.  TODO: Find out.
	ACTIVE Status = "ACTIVE"
)

// GroupPage is the page returned by a pager when traversing over a collection
// of Auto Scale groups.
type GroupPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no Group results.
func (page GroupPage) IsEmpty() (bool, error) {
	groups, err := ExtractGroups(page)

	if err != nil {
		return true, err
	}

	return len(groups) == 0, nil
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page GroupPage) NextPageURL() (string, error) {
	var response struct {
		Links []gophercloud.Link `mapstructure:"groups_links"`
	}

	err := mapstructure.Decode(page.Body, &response)

	if err != nil {
		return "", err
	}

	return gophercloud.ExtractNextURL(response.Links)
}

// ExtractGroups interprets the results of a single page from a List() call,
// producing a slice of Groups.
func ExtractGroups(page pagination.Page) ([]Group, error) {
	casted := page.(GroupPage).Body

	var response struct {
		Groups []Group `mapstructure:"groups"`
	}

	err := mapstructure.Decode(casted, &response)

	if err != nil {
		return nil, err
	}

	return response.Groups, err
}
