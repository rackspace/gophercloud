// +build fixtures

package groups

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rackspace/gophercloud"
	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

// GroupListBody contains the canned body of a groups.List response.
const GroupListBody = `
{
  "groups_links": [],
  "groups": [
    {
      "state": {
        "status": "ACTIVE",
        "desiredCapacity": 2,
        "paused": false,
        "active": [
          {
            "id": "449cead0-48b2-44fe-9107-dea7cdb6d925",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/449cead0-48b2-44fe-9107-dea7cdb6d925",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/449cead0-48b2-44fe-9107-dea7cdb6d925",
                "rel": "bookmark"
              }
            ]
          },
          {
            "id": "d8c2696f-1936-45c7-892d-f5f741ef0f60",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/d8c2696f-1936-45c7-892d-f5f741ef0f60",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/d8c2696f-1936-45c7-892d-f5f741ef0f60",
                "rel": "bookmark"
              }
            ]
          }
        ],
        "pendingCapacity": 0,
        "activeCapacity": 2,
        "name": "first-group"
      },
      "id": "10eb3219-1b12-4b34-b1e4-e10ee4f24c65",
      "links": [
        {
          "href": "https://dfw.autoscale.api.rackspacecloud.com/v1.0/123456/groups/10eb3219-1b12-4b34-b1e4-e10ee4f24c65/",
          "rel": "self"
        }
      ]
    },
    {
      "state": {
        "status": "ACTIVE",
        "desiredCapacity": 3,
        "paused": false,
        "active": [
          {
            "id": "6cca7222-8ab5-4361-ac2c-d35eb0b78ab4",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/6cca7222-8ab5-4361-ac2c-d35eb0b78ab4",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/6cca7222-8ab5-4361-ac2c-d35eb0b78ab4",
                "rel": "bookmark"
              }
            ]
          },
          {
            "id": "44764e46-9ab2-48ce-8512-f7691e0cd9d2",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/44764e46-9ab2-48ce-8512-f7691e0cd9d2",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/44764e46-9ab2-48ce-8512-f7691e0cd9d2",
                "rel": "bookmark"
              }
            ]
          },
          {
            "id": "11a31131-9233-4dac-bcab-15ef06f6b939",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/11a31131-9233-4dac-bcab-15ef06f6b939",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/11a31131-9233-4dac-bcab-15ef06f6b939",
                "rel": "bookmark"
              }
            ]
          }
        ],
        "pendingCapacity": 0,
        "activeCapacity": 3,
        "name": "second-group"
      },
      "id": "e21c7d72-2faa-475a-a35c-8c51d9c66e01",
      "links": [
        {
          "href": "https://dfw.autoscale.api.rackspacecloud.com/v1.0/123456/groups/e21c7d72-2faa-475a-a35c-8c51d9c66e01/",
          "rel": "self"
        }
      ]
    },
    {
      "state": {
        "status": "ACTIVE",
        "desiredCapacity": 2,
        "paused": false,
        "active": [
          {
            "id": "f4ff054b-b78c-4123-98f4-7f0e343c64cd",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/f4ff054b-b78c-4123-98f4-7f0e343c64cd",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/f4ff054b-b78c-4123-98f4-7f0e343c64cd",
                "rel": "bookmark"
              }
            ]
          },
          {
            "id": "c89cfdbf-e3fa-419b-844c-70c3e8016268",
            "links": [
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/c89cfdbf-e3fa-419b-844c-70c3e8016268",
                "rel": "self"
              },
              {
                "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/c89cfdbf-e3fa-419b-844c-70c3e8016268",
                "rel": "bookmark"
              }
            ]
          }
        ],
        "pendingCapacity": 0,
        "activeCapacity": 2,
        "name": "third-group"
      },
      "id": "e34fa1e9-d0f4-47c1-9a01-e531204e1f25",
      "links": [
        {
          "href": "https://dfw.autoscale.api.rackspacecloud.com/v1.0/123456/groups/e34fa1e9-d0f4-47c1-9a01-e531204e1f25/",
          "rel": "self"
        }
      ]
    }
  ]
}
`

// FirstGroupStateBody contains the canned body of a groups.GetState response.
// The response corresponds to the state of first result in GroupListBody.
const FirstGroupStateBody = `
{
  "group": {
    "status": "ACTIVE",
    "desiredCapacity": 2,
    "paused": false,
    "active": [
      {
        "id": "449cead0-48b2-44fe-9107-dea7cdb6d925",
        "links": [
          {
            "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/449cead0-48b2-44fe-9107-dea7cdb6d925",
            "rel": "self"
          },
          {
            "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/449cead0-48b2-44fe-9107-dea7cdb6d925",
            "rel": "bookmark"
          }
        ]
      },
      {
        "id": "d8c2696f-1936-45c7-892d-f5f741ef0f60",
        "links": [
          {
            "href": "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/d8c2696f-1936-45c7-892d-f5f741ef0f60",
            "rel": "self"
          },
          {
            "href": "https://dfw.servers.api.rackspacecloud.com/123456/servers/d8c2696f-1936-45c7-892d-f5f741ef0f60",
            "rel": "bookmark"
          }
        ]
      }
    ],
    "pendingCapacity": 0,
    "activeCapacity": 2,
    "name": "first-group"
  }
}
`

var (
	// FirstGroupState is a State struct corresponding to the state of
	// the first result in GroupListBody.
	FirstGroupState = State{
		Name:            "first-group",
		Status:          ACTIVE,
		DesiredCapacity: 2,
		PendingCapacity: 0,
		ActiveCapacity:  2,
		Paused:          false,
		Errors:          nil,
		Active: []ActiveServer{
			ActiveServer{
				ID: "449cead0-48b2-44fe-9107-dea7cdb6d925",
				Links: []gophercloud.Link{
					gophercloud.Link{
						Href: "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/449cead0-48b2-44fe-9107-dea7cdb6d925",
						Rel:  "self",
					},
					gophercloud.Link{
						Href: "https://dfw.servers.api.rackspacecloud.com/123456/servers/449cead0-48b2-44fe-9107-dea7cdb6d925",
						Rel:  "bookmark",
					},
				},
			},
			ActiveServer{
				ID: "d8c2696f-1936-45c7-892d-f5f741ef0f60",
				Links: []gophercloud.Link{
					gophercloud.Link{
						Href: "https://dfw.servers.api.rackspacecloud.com/v2/123456/servers/d8c2696f-1936-45c7-892d-f5f741ef0f60",
						Rel:  "self",
					},
					gophercloud.Link{
						Href: "https://dfw.servers.api.rackspacecloud.com/123456/servers/d8c2696f-1936-45c7-892d-f5f741ef0f60",
						Rel:  "bookmark",
					},
				},
			},
		},
	}

	// FirstGroup is a Group struct corresponding to the first result in GroupListBody.
	FirstGroup = Group{
		ID:    "10eb3219-1b12-4b34-b1e4-e10ee4f24c65",
		State: FirstGroupState,
	}
)

// HandleGroupListSuccessfully sets up the test server to respond to a group List request.
func HandleGroupListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")

		r.ParseForm()
		marker := r.Form.Get("marker")

		switch marker {
		case "":
			fmt.Fprintf(w, GroupListBody)
		case "e34fa1e9-d0f4-47c1-9a01-e531204e1f25":
			fmt.Fprintf(w, `{ "servers": [] }`)
		default:
			t.Fatalf("/groups invoked with unexpected marker=[%s]", marker)
		}
	})
}

// HandleGroupGetStateSuccessfully sets up the test server to respond to a group GetState request.
func HandleGroupGetStateSuccessfully(t *testing.T) {
	path := "/groups/10eb3219-1b12-4b34-b1e4-e10ee4f24c65/state"

	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(w, FirstGroupStateBody)
	})
}
