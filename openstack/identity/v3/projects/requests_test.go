package projects

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestCreateSuccessful(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "POST")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		testhelper.TestJSONRequest(t, r, `{
			"project": {
				"domain_id": "1789d1",
				"parent_id": "123c56",
				"enabled": true,
				"name": "Test Group",
				"description": "My new project"
			}
		}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
			"project": {
				"domain_id": "1789d1",
				"parent_id": "123c56",
				"enabled": true,
				"id": "263fd9",
				"name": "Test Group",
				"description": "My new project"
			}
    }`)
	})

	opts := ProjectOpts{
		DomainID:    "1789d1",
		ParentID:    "123c56",
		Enabled:     true,
		Name:        "Test Group",
		Description: "My new project",
	}
	result, err := Create(client.ServiceClient(), opts).Extract()
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	if result.ID != "263fd9" {
		t.Errorf("Project id was unexpected [%s]", result.ID)
	}
	if result.DomainID != "1789d1" {
		t.Errorf("Project domain_id was unexpected [%s]", result.DomainID)
	}
	if result.ParentID != "123c56" {
		t.Errorf("Project parent_id was unexpected [%s]", result.ParentID)
	}
	if !result.Enabled {
		t.Errorf("Project enabled was unexpected [%v]", result.Enabled)
	}
	if result.Name != "Test Group" {
		t.Errorf("Project name was unexpected [%s]", result.Name)
	}
	if result.Description != "My new project" {
		t.Errorf("Project description was unexpected [%s]", result.Description)
	}
}

func TestListSinglePage(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"links": {
					"next": null,
					"previous": null
				},
				"projects": [
					{
						"domain_id": "1789d1",
						"parent_id": "123c56",
						"enabled": true,
						"id": "263fd9",
						"name": "Test Group"
					},
					{
						"domain_id": "1789d1",
						"parent_id": "123c56",
						"enabled": true,
						"id": "50ef01",
						"name": "Build Group"
					}
				]
			}
		`)
	})

	count := 0
	err := List(client.ServiceClient(), ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractProjects(page)
		if err != nil {
			return false, err
		}

		expected := []Project{
			Project{
				DomainID: "1789d1",
				ParentID: "123c56",
				Enabled:  true,
				ID:       "263fd9",
				Name:     "Test Group",
			},
			Project{
				DomainID: "1789d1",
				ParentID: "123c56",
				Enabled:  true,
				ID:       "50ef01",
				Name:     "Build Group",
			},
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %#v, got %#v", expected, actual)
		}

		return true, nil
	})
	if err != nil {
		t.Errorf("Unexpected error while paging: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetSuccessful(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects/263fd9", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"project": {
					"domain_id": "1789d1",
					"parent_id": "123c56",
					"enabled": true,
					"id": "263fd9",
					"name": "Test Group",
					"description": "My new project"
				}
			}
		`)
	})

	result, err := Get(client.ServiceClient(), "263fd9").Extract()
	if err != nil {
		t.Fatalf("Error fetching service information: %v", err)
	}
	if result.ID != "263fd9" {
		t.Errorf("Project id was unexpected [%s]", result.ID)
	}
	if result.DomainID != "1789d1" {
		t.Errorf("Project domain_id was unexpected [%s]", result.DomainID)
	}
	if result.ParentID != "123c56" {
		t.Errorf("Project parent_id was unexpected [%s]", result.ParentID)
	}
	if !result.Enabled {
		t.Errorf("Project enabled was unexpected [%v]", result.Enabled)
	}
	if result.Name != "Test Group" {
		t.Errorf("Project name was unexpected [%s]", result.Name)
	}
	if result.Description != "My new project" {
		t.Errorf("Project description was unexpected [%s]", result.Description)
	}
}

func TestUpdateSuccessful(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects/263fd9", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "PUT")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		testhelper.TestJSONRequest(t, r, `{
			"project": {
				"domain_id": "1789d1",
				"parent_id": "123c56",
				"enabled": true,
				"name": "Test Group",
				"description": "My new project"
			}
		}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
			"project": {
				"domain_id": "1789d1",
				"parent_id": "123c56",
				"enabled": true,
				"id": "263fd9",
				"name": "Test Group",
				"description": "My new project"
			}
    }`)
	})

	opts := ProjectOpts{
		DomainID:    "1789d1",
		ParentID:    "123c56",
		Enabled:     true,
		Name:        "Test Group",
		Description: "My new project",
	}
	result, err := Update(client.ServiceClient(), "263fd9", opts).Extract()
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	if result.ID != "263fd9" {
		t.Errorf("Project id was unexpected [%s]", result.ID)
	}
	if result.DomainID != "1789d1" {
		t.Errorf("Project domain_id was unexpected [%s]", result.DomainID)
	}
	if result.ParentID != "123c56" {
		t.Errorf("Project parent_id was unexpected [%s]", result.ParentID)
	}
	if !result.Enabled {
		t.Errorf("Project enabled was unexpected [%v]", result.Enabled)
	}
	if result.Name != "Test Group" {
		t.Errorf("Project name was unexpected [%s]", result.Name)
	}
	if result.Description != "My new project" {
		t.Errorf("Project description was unexpected [%s]", result.Description)
	}
}

func TestDeleteSuccessful(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects/263fd9", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "DELETE")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	res := Delete(client.ServiceClient(), "263fd9")
	testhelper.AssertNoErr(t, res.Err)
}
