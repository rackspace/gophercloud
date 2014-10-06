package extensions

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// GetResult temporarility stores the result of a Get call.
// Use its Extract() method to interpret it as an Extension.
type GetResult struct {
	gophercloud.CommonResult
}

// Extract interprets a GetResult as an Extension.
func (r GetResult) Extract() (*Extension, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Extension *Extension `json:"extension"`
	}

	err := mapstructure.Decode(r.Resp, &res)
	if err != nil {
		return nil, fmt.Errorf("Error decoding OpenStack extension: %v", err)
	}

	return res.Extension, nil
}

// Extension is a struct that represents an OpenStack extension.
type Extension struct {
	Updated     string        `json:"updated"`
	Name        string        `json:"name"`
	Links       []interface{} `json:"links"`
	Namespace   string        `json:"namespace"`
	Alias       string        `json:"alias"`
	Description string        `json:"description"`
}

// ExtensionPage is the page returned by a pager when traversing over a collection of extensions.
type ExtensionPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether an ExtensionPage struct is empty.
func (r ExtensionPage) IsEmpty() (bool, error) {
	is, err := ExtractExtensions(r)
	if err != nil {
		return true, err
	}
	return len(is) == 0, nil
}

// ExtractExtensions accepts a Page struct, specifically an ExtensionPage struct, and extracts the
// elements into a slice of Extension structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractExtensions(page pagination.Page) ([]Extension, error) {
	var resp struct {
		Extensions []Extension `mapstructure:"extensions"`
	}

	err := mapstructure.Decode(page.(ExtensionPage).Body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Extensions, nil
}