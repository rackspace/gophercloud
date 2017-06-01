package snapshots

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

// Snapshot contains all the information associated with an OpenStack Snapshot.
type Snapshot struct {
	// The date when this snapshot was created.
	CreatedAt string `mapstructure:"created_at"`

	// The date when this snapshot was last updated.
	UpdatedAt string `mapstructure:"updated_at"`

	// Human-readable description for the snapshot.
	Description string `mapstructure:"description"`

	// Human-readable display name for the snapshot.
	Name string `mapstructure:"name"`

	// The ID of the snapshots parent snapshot
	SourceVolID string `mapstructure:"volume_id"`

	// Current status of the snapshot.
	Status string `mapstructure:"status"`

	// Arbitrary key-value pairs defined by the user.
	Metadata map[string]string `mapstructure:"metadata"`

	// Unique identifier for the snapshot.
	ID string `mapstructure:"id"`

	// Size of the snapshot in GB.
	Size int `mapstructure:"size"`
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ListResult is a pagination.pager that is returned from a call to the List function.
type ListResult struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no Snapshots.
func (r ListResult) IsEmpty() (bool, error) {
	snapshots, err := ExtractSnapshots(r)
	if err != nil {
		return true, err
	}
	return len(snapshots) == 0, nil
}

// ExtractSnapshots extracts and returns Snapshots. It is used while iterating
// over a snapshots.List call.
func ExtractSnapshots(page pagination.Page) ([]Snapshot, error) {
	var response struct {
		Snapshots []Snapshot `json:"snapshots"`
	}

	err := mapstructure.Decode(page.(ListResult).Body, &response)
	return response.Snapshots, err
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Snapshot object out of the commonResult object.
func (r commonResult) Extract() (*Snapshot, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Snapshot *Snapshot `json:"snapshots"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return res.Snapshot, err
}
