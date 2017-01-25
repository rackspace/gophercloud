package groups

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

	testhelper.Mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"links": {
					"self": "http://example.com/identity/v3/groups",
					"previous": null,
					"next": null
				},
				"groups": [
					{
						"description": "non-admin group",
						"id": "96372bbb152f475aa37e9a76a25a029c",
						"links": {
							"self": "http://example.com/identity/v3/groups/96372bbb152f475aa37e9a76a25a029c"
						},
						"name": "nonadmins",
						"domain_id": "default"
					}
				]
			}
		`)
	})

	count := 0
	err := List(client.ServiceClient(), ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractGroups(page)
		if err != nil {
			return false, err
		}

		expected := []Group{
			Group{
				Description: "non-admin group",
				Name:        "nonadmins",
				ID:          "96372bbb152f475aa37e9a76a25a029c",
				Links:       Link{Self: "http://example.com/identity/v3/groups/96372bbb152f475aa37e9a76a25a029c"},
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
