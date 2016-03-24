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
type State struct {
	// The name of the scaling group.
	Name string `mapstructure:"name" json:"name"`

	// Status of the scaling group.
	Status Status `mapstructure:"status" json:"status"`

	// Number of servers desired.
	DesiredCapacity int `mapstructure:"desiredCapacity" json:"desiredCapacity"`

	// Number of servers in a "BUILDING" state.
	PendingCapacity int `mapstructure:"pendingCapacity" json:"pendingCapacity"`

	// Number of active servers.
	ActiveCapacity int `mapstructure:"activeCapacity" json:"activeCapacity"`

	// Whether a group is paused. All scaling operations are suspended while a
	// group is pasued.
	Paused bool `mapstructure:"paused" json:"paused"`

	// List of active servers. Includes server ID and links.
	Active []ActiveServer `mapstructure:"active" json:"active"`

	// List of errors with human readable messages when a group is in the
	// "ERROR" state.
	Errors []Error `mapstructure:"errors" json:"errors"`
}

// ActiveServer represents an active member server of a scaling group.
type ActiveServer struct {
	// The UUID of the server.
	ID string `mapstructure:"id" json:"id"`

	// Links associated with the server.
	Links []gophercloud.Link `mapstructure:"links" json:"links"`
}

// Error represents a human readable error for groups in an ERROR state.
type Error struct {
	Message string `mapstructure:"message" json:"message"`
}

func (e *Error) Error() string { return e.Message }

// Status indicates the status of an Auto Scale group.
type Status string

// Possible group states.
const (
	ACTIVE   Status = "ACTIVE"
	ERROR    Status = "ERROR"
	DELETING Status = "DELETING"
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
