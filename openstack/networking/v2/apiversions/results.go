package apiversions

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud/pagination"
)

// APIVersion represents an API version for Neutron. It contains the status of
// the API, and its unique ID.
type APIVersion struct {
	Status string `mapstructure:"status" json:"status"`
	ID     string `mapstructure:"id" json:"id"`
}

// APIVersionPage is the page returned by a pager when traversing over a
// collection of API versions.
type APIVersionPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether the page is empty.
func (r APIVersionPage) IsEmpty() (bool, error) {
	is, err := ExtractAPIVersions(r)
	if err != nil {
		return true, err
	}
	return len(is) == 0, nil
}

// ExtractAPIVersion takes a collection page, extracts all of the elements,
// and returns them a slice of APIVersion structs. It is effectively a cast.
func ExtractAPIVersions(page pagination.Page) ([]APIVersion, error) {
	var resp struct {
		Versions []APIVersion `mapstructure:"versions"`
	}

	err := mapstructure.Decode(page.(APIVersionPage).Body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Versions, nil
}

// APIVersionResource represents a generic API resource. It contains the name
// of the resource and its plural collection name.
type APIVersionResource struct {
	Name       string `mapstructure:"name" json:"name"`
	Collection string `mapstructure:"collection" json:"collection"`
}

type APIVersionResourcePage struct {
	pagination.SinglePageBase
}

func (r APIVersionResourcePage) IsEmpty() (bool, error) {
	is, err := ExtractVersionResources(r)
	if err != nil {
		return true, err
	}
	return len(is) == 0, nil
}

func ExtractVersionResources(page pagination.Page) ([]APIVersionResource, error) {
	var resp struct {
		APIVersionResources []APIVersionResource `mapstructure:"resources"`
	}

	err := mapstructure.Decode(page.(APIVersionResourcePage).Body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.APIVersionResources, nil
}