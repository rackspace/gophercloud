package snapshots

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the Snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	// The snapshot description [OPTIONAL]
	Description string

	// One or more metadata key and value pairs to associate with the snapshot [OPTIONAL]
	Metadata map[string]string

	// The snapshot name [OPTIONAL]
	Name string

	// the ID of the source volume  to snapshot
	VolumeID string

	//Force creation even if volume is attached
	Force bool
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	s := make(map[string]interface{})

	if opts.VolumeID == "" {
		return nil, fmt.Errorf("Required CreateOpts field 'VolumeID' not set.")
	}
	s["volume_id"] = opts.VolumeID

	if opts.Description != "" {
		s["description"] = opts.Description
	}
	if opts.Metadata != nil {
		s["metadata"] = opts.Metadata
	}
	if opts.Name != "" {
		s["name"] = opts.Name
	}

	s["force"] = opts.Force

	return map[string]interface{}{"snapshot": s}, nil
}

// Create will create a new Snapshot based on the values in CreateOpts. To extract
// the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToSnapshotCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(createURL(client), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return res
}

// Delete will delete the existing Snapshot with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var res DeleteResult
	_, res.Err = client.Delete(deleteURL(client, id), nil)
	return res
}

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot object
// from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = client.Get(getURL(client, id), &res.Body, nil)
	return res
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ListOpts holds options for listing Snapshots. It is passed to the Snapshots.List
// function.
type ListOpts struct {
	// admin-only option. Set it to true to see all tenant snapshots.
	AllTenants bool `q:"all_tenants"`
	// List only snapshots that contain Metadata.
	Metadata map[string]string `q:"metadata"`
	// List only snapshots that have Name as the display name.
	Name string `q:"name"`
	// List only snapshots that have a status of Status.
	Status string `q:"status"`
}

// ToSnapshotListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSnapshotListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns Snapshots optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToSnapshotListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	createPage := func(r pagination.PageResult) pagination.Page {
		return ListResult{pagination.SinglePageBase(r)}
	}

	return pagination.NewPager(client, url, createPage)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSnapshotUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Snapshot. This object is passed
// to the Snapshots.Update function. For more information about the parameters, see
// the Snapshot object.
type UpdateOpts struct {
	// OPTIONAL
	Name string
	// OPTIONAL
	Description string
}

// ToSnapshotUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToSnapshotUpdateMap() (map[string]interface{}, error) {
	s := make(map[string]interface{})

	if opts.Description != "" {
		s["description"] = opts.Description
	}
	if opts.Name != "" {
		s["name"] = opts.Name
	}
	return map[string]interface{}{"snapshot": s}, nil
}

// Update will update the Snapshot with provided information. To extract the updated
// Snapshot from the response, call the Extract method on the UpdateResult.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) UpdateResult {
	var res UpdateResult

	reqBody, err := opts.ToSnapshotUpdateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Put(updateURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

// IDFromName is a convienience function that returns a server's ID given its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	snapshotCount := 0
	snapshotID := ""
	if name == "" {
		return "", fmt.Errorf("A snapshot name must be provided.")
	}
	pager := List(client, nil)
	pager.EachPage(func(page pagination.Page) (bool, error) {
		snapshotList, err := ExtractSnapshots(page)
		if err != nil {
			return false, err
		}

		for _, s := range snapshotList {
			if s.Name == name {
				snapshotCount++
				snapshotID = s.ID
			}
		}
		return true, nil
	})

	switch snapshotCount {
	case 0:
		return "", fmt.Errorf("Unable to find snapshot: %s", name)
	case 1:
		return snapshotID, nil
	default:
		return "", fmt.Errorf("Found %d snapshots matching %s", snapshotCount, name)
	}
}
