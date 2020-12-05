package domains

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

type commonResult struct {
	gophercloud.Result
}

// Link the object to hold a project link.
type Link struct {
	Self string `json:"self,omitempty"`
}

// Domain is main struct for holding domain attributes.
type Domain struct {
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Links       Link   `json:"links"`
}

// PairResult the object to error for failed pairs.
type PairResult struct {
	commonResult
}

// DomainPage is a single page of Domain results.
type DomainPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p DomainPage) IsEmpty() (bool, error) {
	domains, err := ExtractDomains(p)
	if err != nil {
		return true, err
	}
	return len(domains) == 0, nil
}

// ExtractDomains extracts a slice of Domains from a Collection acquired from List.
func ExtractDomains(page pagination.Page) ([]Domain, error) {
	var response struct {
		Domains []Domain `mapstructure:"domains"`
	}

	err := mapstructure.Decode(page.(DomainPage).Body, &response)
	return response.Domains, err
}
