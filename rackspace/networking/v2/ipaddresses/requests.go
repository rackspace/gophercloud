package sharedips

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSharedIPListQuery() (string, error)
}

// ListOpts are options for listing shared IPs. Though currently empty, having
// it lets us add fields later should they get added to the API.
type ListOpts struct{}

// ToSharedIPListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSharedIPListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List returns a Pager which allows you to iterate over a collection of
// shared IPs.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToSharedIPListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return IPAddressPage{pagination.SinglePageBase(r)}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSharedIPCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for provisioning a shared IP address. This object is passed to
// the sharedips.Create function.
type CreateOpts struct {
	// REQUIRED
	NetworkID string
	// REQUIRED
	PortIDs []string
	// REQUIRED
	Version gophercloud.IPVersion
	// REQUIRED
	TenantID string
}

// ToSharedIPCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSharedIPCreateMap() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	if opts.NetworkID == "" {
		return nil, fmt.Errorf("Required CreateOpts field 'NetworkID' not set.")
	}
	v["network_id"] = opts.NetworkID

	if len(opts.PortIDs) < 2 {
		return nil, fmt.Errorf("Required CreateOpts field 'PortIDs' must contain at least 2 port IDs.")
	}
	v["port_ids"] = opts.PortIDs

	if opts.Version != 4 && opts.Version != 6 {
		return nil, fmt.Errorf("Required CreateOpts field 'Version' not set.")
	}
	v["version"] = opts.Version

	if opts.TenantID == "" {
		return nil, fmt.Errorf("Required CreateOpts field 'TenantID' not set.")
	}
	v["tenant_id"] = opts.TenantID

	return map[string]interface{}{"ip_address": v}, nil
}

// Create will provision a new shared IP address based on the values in CreateOpts. To extract
// the IPAddress object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToSharedIPCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(createURL(client), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return res
}

// Delete will deallocate the existing shared IP with the provided ID from the tenant.
// Before using this operation, all IP associations must be removed from the IP address
// by using the Disassociate function.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var res DeleteResult
	_, res.Err = client.Delete(deleteURL(client, id), nil)
	return res
}

// Get retrieves the shared IP with the provided ID. To extract the IPAddress object
// from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = client.Get(getURL(client, id), &res.Body, nil)
	return res
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSharedIPUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing shared IP. This object is passed
// to the sharedips.Update function.
type UpdateOpts struct {
	// REQUIRED
	PortIDs []string
}

// ToSharedIPUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToSharedIPUpdateMap() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	if len(opts.PortIDs) < 2 {
		return nil, fmt.Errorf("Required CreateOpts field 'PortIDs' must contain at least 2 port IDs.")
	}
	v["port_ids"] = opts.PortIDs

	return map[string]interface{}{"ip_address": v}, nil
}

// Update will update the shared IP with provided information. To extract the updated
// IPAddress from the response, call the Extract method on the UpdateResult.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) UpdateResult {
	var res UpdateResult

	reqBody, err := opts.ToSharedIPUpdateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Put(updateURL(client, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

// ListByServerOptsBuilder allows extensions to add additional parameters to the
// ListByServer request.
type ListByServerOptsBuilder interface {
	ToSharedIPListByServerQuery() (string, error)
}

// ListByServerOpts are options for listing the shared IPs for a server. Though currently empty, having
// it lets us add fields later should they get added to the API.
type ListByServerOpts struct{}

// ToSharedIPListByServerQuery formats a ListByServerOpts into a query string.
func (opts ListOpts) ToSharedIPListByServerQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// ListByServer returns a Pager which allows you to iterate over a collection of
// IPAssociations.
func ListByServer(c *gophercloud.ServiceClient, serverID string, opts ListByServerOptsBuilder) pagination.Pager {
	url := listByServerURL(c, serverID)
	if opts != nil {
		query, err := opts.ToSharedIPListByServerQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return IPAssociationPage{pagination.SinglePageBase(r)}
	})
}

// Associate will associate the server with a shared IP address.
func Associate(client *gophercloud.ServiceClient, serverID, sharedIPID string) AssociateResult {
	var res AssociateResult

	_, res.Err = client.Post(associateURL(client, serverID, sharedIPID), nil, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return res
}

// Disassociate will delete the association between the server and the IP address.
func Disassociate(client *gophercloud.ServiceClient, serverID, sharedIPID string) DisassociateResult {
	var res DisassociateResult
	_, res.Err = client.Delete(disassociateURL(client, serverID, sharedIPID), nil)
	return res
}

// GetByServer retrieves the shared IP with the provided ID and server ID. To extract
// the IPAssociation object from the response, call the Extract method on the GetByServerResult.
func GetByServer(client *gophercloud.ServiceClient, serverID, sharedIPID string) GetByServerResult {
	var res GetByServerResult
	_, res.Err = client.Get(getByServerURL(client, serverID, sharedIPID), &res.Body, nil)
	return res
}
