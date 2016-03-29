// +build fixtures

package launch

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

// LaunchGetBody contains the canned body of a launch.Get response.
const LaunchGetBody = `
{
  "launchConfiguration": {
    "args": {
      "loadBalancers": [
        {
          "port": 443,
          "loadBalancerId": 123456
        }
      ],
      "server": {
        "name": "server",
        "imageRef": "40155f16-21d4-4ac1-ad65-c409d94b8c7c",
        "key_name": "gophercloud",
        "flavorRef": "performance1-4",
        "user_data": "thequickbrownfoxjumpsoverthelazydog",
        "config_drive": true,
        "networks": [
          {"uuid": "00000000-0000-0000-0000-000000000000"},
          {"uuid": "11111111-1111-1111-1111-111111111111"}
        ],
        "metadata": {
          "foo": "bar"
        }
      }
    },
    "type": "launch_server"
  }
}
`

// LaunchConfig is a Configuration struct corresponding to the result in LaunchGetBody.
var LaunchConfig = Configuration{
	Type: "launch_server",

	Args: Args{
		Server: map[string]interface{}{
			"name":         "server",
			"imageRef":     "40155f16-21d4-4ac1-ad65-c409d94b8c7c",
			"key_name":     "gophercloud",
			"flavorRef":    "performance1-4",
			"user_data":    "thequickbrownfoxjumpsoverthelazydog",
			"config_drive": true,
			"networks": []map[string]string{
				{"uuid": "00000000-0000-0000-0000-000000000000"},
				{"uuid": "11111111-1111-1111-1111-111111111111"},
			},
			"metadata": map[string]string{
				"foo": "bar",
			},
		},
		LoadBalancers: []map[string]interface{}{
			{
				"port":           443,
				"loadBalancerId": 123456,
			},
		},
		DrainingTimeout: 0, // TODO: This is optional, is this right?
	},
}

// HandleLanuchGetSuccessfully sets up the test server to respond to a launch Gist request.
func HandleLanuchGetSuccessfully(t *testing.T) {
	path := "/groups/10eb3219-1b12-4b34-b1e4-e10ee4f24c65/launch"

	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(w, LaunchGetBody)
	})

}
