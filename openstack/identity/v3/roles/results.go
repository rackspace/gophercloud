package roles

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

// RoleAssignment is the result of a role assignments query.
type RoleAssignment struct {
	Role  Role  `json:"role,omitempty"`
	Scope Scope `json:"scope,omitempty"`
	User  User  `json:"user,omitempty"`
	Group Group `json:"group,omitempty"`
}

// Link the object to hold a project link.
type Link struct {
	Self string `json:"self,omitempty"`
}

type Role struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	DomainID string `json:"domain_id,omitempty"`
	Links    Link   `json:"links"`
}

type Scope struct {
	Domain  Domain  `json:"domain,omitempty"`
	Project Project `json:"domain,omitempty"`
}

type Domain struct {
	ID string `json:"id,omitempty"`
}

type Project struct {
	ID string `json:"id,omitempty"`
}

type User struct {
	ID string `json:"id,omitempty"`
}

type Group struct {
	ID string `json:"id,omitempty"`
}

// RoleAssignmentsPage is a single page of RoleAssignments results.
type RoleAssignmentsPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p RoleAssignmentsPage) IsEmpty() (bool, error) {
	roleAssignments, err := ExtractRoleAssignments(p)
	if err != nil {
		return true, err
	}
	return len(roleAssignments) == 0, nil
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page RoleAssignmentsPage) NextPageURL() (string, error) {
	type resp struct {
		Links struct {
			Next string `mapstructure:"next"`
		} `mapstructure:"links"`
	}

	var r resp
	err := mapstructure.Decode(page.Body, &r)
	if err != nil {
		return "", err
	}

	return r.Links.Next, nil
}

// ExtractRoleAssignments extracts a slice of RoleAssignments from a Collection acquired from List.
func ExtractRoleAssignments(page pagination.Page) ([]RoleAssignment, error) {
	var response struct {
		RoleAssignments []RoleAssignment `mapstructure:"role_assignments"`
	}

	err := mapstructure.Decode(page.(RoleAssignmentsPage).Body, &response)
	return response.RoleAssignments, err
}

type response struct {
	Role Role `json:"role"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a CreateResult as a concrete Service.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Role, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Role `json:"role"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return &res.Role, err
}

// CreateResult the object to hold a role link.
type CreateResult struct {
	commonResult
}

// RolePage is a single page of Role results.
type RolePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p RolePage) IsEmpty() (bool, error) {
	roles, err := ExtractRoles(p)
	if err != nil {
		return true, err
	}
	return len(roles) == 0, nil
}

// ExtractRoles extracts a slice of Roles from a Collection acquired from List.
func ExtractRoles(page pagination.Page) ([]Role, error) {
	var response struct {
		Roles []Role `mapstructure:"roles"`
	}

	err := mapstructure.Decode(page.(RolePage).Body, &response)
	return response.Roles, err
}
