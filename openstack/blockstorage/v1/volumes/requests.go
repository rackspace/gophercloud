package volumes

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Volume. This object is passed to
// the volumes.Create function. For more information about these parameters,
// see the Volume object.
type CreateOpts struct {
	// OPTIONAL
	Availability string
	// OPTIONAL
	Description string
	// OPTIONAL
	Metadata map[string]string
	// OPTIONAL
	Name string
	// REQUIRED
	Size int
	// OPTIONAL
	SnapshotID, SourceVolID, ImageID string
	// OPTIONAL
	VolumeType string
}

// ToVolumeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	if opts.Size == 0 {
		return nil, fmt.Errorf("Required CreateOpts field 'Size' not set.")
	}
	v["size"] = opts.Size

	if opts.Availability != "" {
		v["availability_zone"] = opts.Availability
	}
	if opts.Description != "" {
		v["display_description"] = opts.Description
	}
	if opts.ImageID != "" {
		v["imageRef"] = opts.ImageID
	}
	if opts.Metadata != nil {
		v["metadata"] = opts.Metadata
	}
	if opts.Name != "" {
		v["display_name"] = opts.Name
	}
	if opts.SourceVolID != "" {
		v["source_volid"] = opts.SourceVolID
	}
	if opts.SnapshotID != "" {
		v["snapshot_id"] = opts.SnapshotID
	}
	if opts.VolumeType != "" {
		v["volume_type"] = opts.VolumeType
	}

	return map[string]interface{}{"volume": v}, nil
}

// Create will create a new Volume based on the values in CreateOpts. To extract
// the Volume object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToVolumeCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(createURL(client), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return res
}

// Delete will delete the existing Volume with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var res DeleteResult
	_, res.Err = client.Delete(deleteURL(client, id), nil)
	return res
}

// Get retrieves the Volume with the provided ID. To extract the Volume object
// from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = client.Get(getURL(client, id), &res.Body, nil)
	return res
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeListQuery() (string, error)
}

// ListOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListOpts struct {
	// admin-only option. Set it to true to see all tenant volumes.
	AllTenants bool `q:"all_tenants"`
	// List only volumes that contain Metadata.
	Metadata map[string]string `q:"metadata"`
	// List only volumes that have Name as the display name.
	Name string `q:"name"`
	// List only volumes that have a status of Status.
	Status string `q:"status"`
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
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
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Volume. This object is passed
// to the volumes.Update function. For more information about the parameters, see
// the Volume object.
type UpdateOpts struct {
	// OPTIONAL
	Name string
	// OPTIONAL
	Description string
	// OPTIONAL
	Metadata map[string]string
}

// ToVolumeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	if opts.Description != "" {
		v["display_description"] = opts.Description
	}
	if opts.Metadata != nil {
		v["metadata"] = opts.Metadata
	}
	if opts.Name != "" {
		v["display_name"] = opts.Name
	}

	return map[string]interface{}{"volume": v}, nil
}

// Update will update the Volume with provided information. To extract the updated
// Volume from the response, call the Extract method on the UpdateResult.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) UpdateResult {
	var res UpdateResult

	reqBody, err := opts.ToVolumeUpdateMap()
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
	volumeCount := 0
	volumeID := ""
	if name == "" {
		return "", fmt.Errorf("A volume name must be provided.")
	}
	pager := List(client, nil)
	pager.EachPage(func(page pagination.Page) (bool, error) {
		volumeList, err := ExtractVolumes(page)
		if err != nil {
			return false, err
		}

		for _, s := range volumeList {
			if s.Name == name {
				volumeCount++
				volumeID = s.ID
			}
		}
		return true, nil
	})

	switch volumeCount {
	case 0:
		return "", fmt.Errorf("Unable to find volume: %s", name)
	case 1:
		return volumeID, nil
	default:
		return "", fmt.Errorf("Found %d volumes matching %s", volumeCount, name)
	}
}

type AttachOptsBuilder interface {
	ToVolumeAttachMap() (map[string]interface{}, error)
}

type AttachOpts struct {
	MountPoint 		string
	InstanceUUID 	string
	HostName		string
	Mode			string
}

func (opts AttachOpts) ToVolumeAttachMap() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	if opts.MountPoint != "" {
		v["mountpoint"] = opts.MountPoint
	}
	if opts.Mode != "" {
		v["mode"] = opts.Mode
	}
	if opts.InstanceUUID != "" {
		v["instance_uuid"] = opts.InstanceUUID
	}
	if opts.HostName != "" {
		v["host_name"] = opts.HostName
	}

	return map[string]interface{}{"os-attach": v}, nil
}

func Attach(client *gophercloud.ServiceClient, id string, opts AttachOptsBuilder) AttachResult {
	var res AttachResult

	reqBody, err := opts.ToVolumeAttachMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(attachURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return res
}

func Detach(client *gophercloud.ServiceClient, id string) DetachResult {
	var res DetachResult

	v := make(map[string]interface{})
	reqBody := map[string]interface{}{"os-detach": v}

	_, res.Err = client.Post(detachURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return res
}

func Reserve(client *gophercloud.ServiceClient, id string) ReserveResult {
	var res ReserveResult

	v := make(map[string]interface{})
	reqBody := map[string]interface{}{"os-reserve": v}

	_, res.Err = client.Post(reserveURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return res
}

func Unreserve(client *gophercloud.ServiceClient, id string) UnreserveResult {
	var res UnreserveResult

	v := make(map[string]interface{})
	reqBody := map[string]interface{}{"os-unreserve": v}

	_, res.Err = client.Post(unreserveURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return res
}

type ConnectorOptsBuilder interface {
	ToConnectorMap() (map[string]interface{}, error)
}

type ConnectorOpts struct {
	IP 			string
	Host 		string
	Initiator	string
	Wwpns		string
	Wwnns 		string
	Multipath 	bool
	Platform 	string
	OSType		string
}

func (opts ConnectorOpts) ToConnectorMap() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	if opts.IP != "" {
		v["ip"] = opts.IP
	}
	if opts.Host != "" {
		v["host"] = opts.Host
	}
	if opts.Initiator != "" {
		v["initiator"] = opts.Initiator
	}
	if opts.Wwpns != "" {
		v["wwpns"] = opts.Wwpns
	}
	if opts.Wwnns != "" {
		v["wwnns"] = opts.Wwnns
	}

	v["multipath"] = opts.Multipath

	if opts.Platform != "" {
		v["platform"] = opts.Platform
	}
	if opts.OSType != "" {
		v["os_type"] = opts.OSType
	}

	return map[string]interface{}{"connector": v}, nil
}

func InitializeConnection(client *gophercloud.ServiceClient, id string, opts ConnectorOpts) InitializeConnectionResult {
	var res InitializeConnectionResult

	connctorMap, err := opts.ToConnectorMap()
	if err != nil {
		res.Err = err
		return res
	}

	reqBody := map[string]interface{}{"os-initialize_connection": connctorMap}

	_, res.Err = client.Post(attachURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return res
}

func TerminateConnection(client *gophercloud.ServiceClient, id string, opts ConnectorOpts) TerminateConnectionResult {
	var res TerminateConnectionResult

	connctorMap, err := opts.ToConnectorMap()
	if err != nil {
		res.Err = err
		return res
	}

	reqBody := map[string]interface{}{"os-terminate_connection": connctorMap}

	_, res.Err = client.Post(attachURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	return res
}