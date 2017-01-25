package roles

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestListSinglePageRA(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/role_assignments", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
                "role_assignments": [
                    {
                        "links": {
                            "assignment": "http://identity:35357/v3/domains/161718/users/313233/roles/123456"
                        },
                        "role": {
                            "id": "123456"
                        },
                        "scope": {
                            "domain": {
                                "id": "161718"
                            }
                        },
                        "user": {
                            "id": "313233"
                        }
                    },
                    {
                        "links": {
                            "assignment": "http://identity:35357/v3/projects/456789/groups/101112/roles/123456",
                            "membership": "http://identity:35357/v3/groups/101112/users/313233"
                        },
                        "role": {
                            "id": "123456"
                        },
                        "scope": {
                            "project": {
                                "id": "456789"
                            }
                        },
                        "user": {
                            "id": "313233"
                        }
                    }
                ],
                "links": {
                    "self": "http://identity:35357/v3/role_assignments?effective",
                    "previous": null,
                    "next": null
                }
            }
		`)
	})

	count := 0
	err := ListAssignments(client.ServiceClient(), ListAssignmentsOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractRoleAssignments(page)
		if err != nil {
			return false, err
		}

		expected := []RoleAssignment{
			RoleAssignment{
				Role:  Role{ID: "123456"},
				Scope: Scope{Domain: Domain{ID: "161718"}},
				User:  User{ID: "313233"},
				Group: Group{},
			},
			RoleAssignment{
				Role:  Role{ID: "123456"},
				Scope: Scope{Project: Project{ID: "456789"}},
				User:  User{ID: "313233"},
				Group: Group{},
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

func TestCreateSuccessful(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "POST")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		testhelper.TestJSONRequest(t, r, `{
			"role": {
				"domain_id": "92e782c4988642d783a95f4a87c3fdd7",
				"name": "developer"
			}
		}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
			"role": {
				"domain_id": "92e782c4988642d783a95f4a87c3fdd7",
				"id": "1e443fa8cee3482a8a2b6954dd5c8f12",
				"links": {
					"self": "http://example.com/identity/v3/roles/1e443fa8cee3482a8a2b6954dd5c8f12"
				},
				"name": "developer"
			}
		}`)
	})

	role := CreateOpts{
		DomainID: "92e782c4988642d783a95f4a87c3fdd7",
		Name:     "developer",
	}

	result, err := Create(client.ServiceClient(), role).Extract()
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	if result.Name != "developer" {
		t.Errorf("Role name was unexpected [%s]", result.Name)
	}
}

func TestListSinglePage(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"links": {
					"next": null,
					"previous": null,
					"self": "http://example.com/identity/v3/roles"
				},
				"roles": [
					{
						"id": "5318e65d75574c17bf5339d3df33a5a3",
						"links": {
							"self": "http://example.com/identity/v3/roles/5318e65d75574c17bf5339d3df33a5a3"
						},
						"name": "admin"
					},
					{
						"id": "9fe2ff9ee4384b1894a90878d3e92bab",
						"links": {
							"self": "http://example.com/identity/v3/roles/9fe2ff9ee4384b1894a90878d3e92bab"
						},
						"name": "_member_"
					}
				]
			}
		`)
	})

	count := 0
	err := List(client.ServiceClient(), ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractRoles(page)
		if err != nil {
			return false, err
		}

		expected := []Role{
			Role{
				ID:          "5318e65d75574c17bf5339d3df33a5a3",
				Name:        "admin",
				Links:       Link{Self: "http://example.com/identity/v3/roles/5318e65d75574c17bf5339d3df33a5a3"},
			},
			Role{
				ID:          "9fe2ff9ee4384b1894a90878d3e92bab",
				Name:        "_member_",
				Links:       Link{Self: "http://example.com/identity/v3/roles/9fe2ff9ee4384b1894a90878d3e92bab"},
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
