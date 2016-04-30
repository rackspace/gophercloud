package tenants

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// Tenant is a grouping of users in the identity service.
type Tenant struct {
	// ID is a unique identifier for this tenant.
	ID string `mapstructure:"id"`

	// Name is a friendlier user-facing name for this tenant.
	Name string `mapstructure:"name"`

	// Description is a human-readable explanation of this Tenant's purpose.
	Description string `mapstructure:"description"`

	// Enabled indicates whether or not a tenant is active.
	Enabled bool `mapstructure:"enabled"`
}

// TenantPage is a single page of Tenant results.
type TenantPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Tenants contains any results.
func (page TenantPage) IsEmpty() (bool, error) {
	tenants, err := ExtractTenants(page)
	if err != nil {
		return false, err
	}
	return len(tenants) == 0, nil
}

// NextPageURL extracts the "next" link from the tenants_links section of the result.
func (page TenantPage) NextPageURL() (string, error) {
	type resp struct {
		Links []gophercloud.Link `mapstructure:"tenants_links"`
	}

	var r resp
	err := mapstructure.Decode(page.Body, &r)
	if err != nil {
		return "", err
	}

	return gophercloud.ExtractNextURL(r.Links)
}

// ExtractTenants returns a slice of Tenants contained in a single page of results.
func ExtractTenants(page pagination.Page) ([]Tenant, error) {
	casted := page.(TenantPage).Body
	var response struct {
		Tenants []Tenant `mapstructure:"tenants"`
	}

	err := mapstructure.Decode(casted, &response)
	return response.Tenants, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a Tenant, if possible.
func (r commonResult) Extract() (*User, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Tenant Tenant `mapstructure:"tenant"`
	}

	err := mapstructure.Decode(r.Body, &response)

	return &response.Tenant, err
}

// CreateResult represents the result of a Create operation
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation
type DeleteResult struct {
	commonResult
}
