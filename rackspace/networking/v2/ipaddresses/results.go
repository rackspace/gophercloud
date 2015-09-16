package sharedips

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// IPAddress represents a shared IP address.
type IPAddress struct {
	ID        string                `mapstructure:"id"`
	NetworkID string                `mapstructure:"network_id"`
	Address   string                `mapstructure:"address"`
	PortIDs   []string              `mapstructure:"port_ids"`
	SubnetID  string                `mapstructure:"subnet_id"`
	TenantID  string                `mapstructure:"tenant_id"`
	Version   gophercloud.IPVersion `mapstructure:"version"`
	Type      string                `mapstructure:"type"`
}

// IPAddressPage is the page returned by a pager when traversing over a
// collection of shared IP addresses.
type IPAddressPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether an IPAddressPage struct is empty.
func (p IPAddressPage) IsEmpty() (bool, error) {
	is, err := ExtractIPAddresses(p)
	if err != nil {
		return true, nil
	}
	return len(is) == 0, nil
}

// ExtractIPAddresses accepts a Page struct, specifically an IPAddressPage struct,
// and extracts the elements into a slice of IPAddress structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractIPAddresses(page pagination.Page) ([]IPAddress, error) {
	var resp struct {
		IPAddresses []IPAddress `mapstructure:"ip_addresses"`
	}
	err := mapstructure.WeakDecode(page.(IPAddressPage).Body, &resp)

	return resp.IPAddresses, err
}

type commonIPAddressResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a shared IP resource.
func (r commonIPAddressResult) Extract() (*IPAddress, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		IPAddress *IPAddress `mapstructure:"ip_address"`
	}

	err := mapstructure.WeakDecode(r.Body, &res)

	return res.IPAddress, err
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	commonIPAddressResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonIPAddressResult
}

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	commonIPAddressResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// IPAssociation represents a shared IP address for a server.
type IPAssociation struct {
	ID      string `mapstructure:"id"`
	Address string `mapstructure:"address"`
}

// IPAssociationPage is the page returned by a pager when traversing over the
// collection of shared IP addresses for a server.
type IPAssociationPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether an IPAssociationPage struct is empty.
func (p IPAssociationPage) IsEmpty() (bool, error) {
	is, err := ExtractIPAssociations(p)
	if err != nil {
		return true, nil
	}
	return len(is) == 0, nil
}

// ExtractIPAssociations accepts a Page struct, specifically an IPAssociationPage struct,
// and extracts the elements into a slice of IPAssociation structs.
func ExtractIPAssociations(page pagination.Page) ([]IPAssociation, error) {
	var resp struct {
		IPAssociations []IPAssociation `mapstructure:"ip_associations"`
	}

	err := mapstructure.Decode(page.(IPAssociationPage).Body, &resp)

	return resp.IPAssociations, err
}

type commonIPAssociationResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a shared IP association.
func (r commonIPAssociationResult) Extract() (*IPAssociation, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		IPAssociation *IPAssociation `mapstructure:"ip_association"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return res.IPAssociation, err
}

// AssociateResult represents the result of an Associate operation.
type AssociateResult struct {
	commonIPAssociationResult
}

// GetByServerResult represents the result of a GetByServer operation.
type GetByServerResult struct {
	commonIPAssociationResult
}

// DisassociateResult represents the result of a Disassociate operation.
type DisassociateResult struct {
	gophercloud.ErrResult
}
