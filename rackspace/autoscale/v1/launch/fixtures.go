// +build fixtures

package launch

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
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
        "user_data": "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQ=",
        "personality": [
          {
            "path": "/etc/motd",
            "contents": "Z29waGVyY2xvdWQ="
          }
        ],
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

// Examples of various components of configuration arguments.
var (
	ExampleUserData = []byte("Lorem ipsum dolor sit amet")

	MOTDFile = servers.File{
		Path:     "/etc/motd",
		Contents: []byte("gophercloud"),
	}

	CloudLB = LoadBalancer{
		CloudLoadBalancerID: 123456,
		Port:                443,
	}
)

// LaunchConfig is a Configuration corresponding to the result in LaunchGetBody.
var LaunchConfig = Configuration{
	Type: LaunchServer,

	Args: Args{
		Server: Server{
			Name:        "server",
			ImageRef:    "40155f16-21d4-4ac1-ad65-c409d94b8c7c",
			KeyName:     "gophercloud",
			FlavorRef:   "performance1-4",
			UserData:    ExampleUserData,
			ConfigDrive: true,
			Personality: servers.Personality{
				&MOTDFile,
			},
			Networks: []Network{
				PublicNet,
				ServiceNet,
			},
			Metadata: map[string]interface{}{
				"foo": "bar",
			},
		},
		LoadBalancers: []LoadBalancer{
			CloudLB,
		},
		DrainingTimeout: 0,
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
