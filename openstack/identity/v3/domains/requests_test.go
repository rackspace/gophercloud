package domains

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestListSinglePage(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
                "domains": [
                    {
                        "description": "Used for swift functional testing",
                        "enabled": true,
                        "id": "5a75994a383c449184053ff7270c4e91",
                        "links": {
                            "self": "http://example.com/identity/v3/domains/5a75994a383c449184053ff7270c4e91"
                        },
                        "name": "swift_test"
                    },
                    {
                        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
                        "enabled": true,
                        "id": "default",
                        "links": {
                            "self": "http://example.com/identity/v3/domains/default"
                        },
                        "name": "Default"
                    }
                ],
                "links": {
                    "next": null,
                    "previous": null,
                    "self": "http://example.com/identity/v3/domains"
                }
            }
		`)
	})

	count := 0
	err := List(client.ServiceClient(), ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractDomains(page)
		if err != nil {
			return false, err
		}

		expected := []Domain{
			Domain{
				Description: "Used for swift functional testing",
				ID:          "5a75994a383c449184053ff7270c4e91",
				Name:        "swift_test",
				Enabled:     true,
				Links:       Link{Self: "http://example.com/identity/v3/domains/5a75994a383c449184053ff7270c4e91"},
			},
			Domain{
				Description: "Owns users and tenants (i.e. projects) available on Identity API v2.",
				ID:          "default",
				Name:        "Default",
				Enabled:     true,
				Links:       Link{Self: "http://example.com/identity/v3/domains/default"},
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

func TestPairDomainGroupAndRole(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/domains/5a75994a383c449184053ff7270c4e91/groups/5a75994a383c449184053ff7270c4e92/roles/5a75994a383c449184053ff7270c4e93", func(w http.ResponseWriter, r *http.Request) {
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
