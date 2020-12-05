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
		    "description": "My new project",
		    "domain_id": "default",
		    "enabled": true,
		    "is_domain": true,
		    "name": "myNewProject"
		  }
		}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
            "project": {
                "description": "My new project",
                "domain_id": "default",
                "enabled": true,
                "is_domain": false,
                "name": "myNewProject",
                "id": "1234567",
                "links": {
                    "self": "http://os.test.com/v3/identity/projects/1234567"
                }
            }
        }`)
	})

	project := CreateOpts{
		IsDomain:    true,
		Description: "My new project",
		DomainID:    "default",
		Enabled:     true,
		Name:        "myNewProject",
	}

	result, err := Create(client.ServiceClient(), project).Extract()
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	if result.Description != "My new project" {
		t.Errorf("Project description was unexpected [%s]", result.Description)
	}
}

func TestPairProjectGroupAndRole(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects/5a75994a383c449184053ff7270c4e91/groups/5a75994a383c449184053ff7270c4e92/roles/5a75994a383c449184053ff7270c4e93", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "PUT")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	pair := PairOpts{
		ID:      "5a75994a383c449184053ff7270c4e91",
		GroupID: "5a75994a383c449184053ff7270c4e92",
		RoleID:  "5a75994a383c449184053ff7270c4e93",
	}

	err := Pair(client.ServiceClient(), pair)
	if err != nil {
		t.Fatalf("Unexpected error from Pair: %v", err)
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
					"previous": null,
					"self": "http://example.com/identity/v3/projects"
				},
				"projects": [
					{
						"is_domain": false,
						"description": null,
						"domain_id": "",
						"enabled": true,
						"id": "0c4e939acacf4376bdcd1129f1a054ad",
						"links": {
							"self": "http://example.com/identity/v3/projects/0c4e939acacf4376bdcd1129f1a054ad"
						},
						"name": "admin",
						"parent_id": null
					},
					{
						"is_domain": false,
						"description": null,
						"domain_id": "",
						"enabled": true,
						"id": "0cbd49cbf76d405d9c86562e1d579bd3",
						"links": {
							"self": "http://example.com/identity/v3/projects/0cbd49cbf76d405d9c86562e1d579bd3"
						},
						"name": "demo",
						"parent_id": null
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
				ID:       "0c4e939acacf4376bdcd1129f1a054ad",
				IsDomain: false,
				DomainID: "",
				Enabled:  true,
				Name:     "admin",
				Links:    Link{Self: "http://example.com/identity/v3/projects/0c4e939acacf4376bdcd1129f1a054ad"},
			},
			Project{
				ID:       "0cbd49cbf76d405d9c86562e1d579bd3",
				IsDomain: false,
				DomainID: "",
				Enabled:  true,
				Name:     "demo",
				Links:    Link{Self: "http://example.com/identity/v3/projects/0cbd49cbf76d405d9c86562e1d579bd3"},
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
