// +build fixtures

package floatingips

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

// ListOutput is a sample response to a List call.
const ListOutput = `
{
  "floating_ips": [
    {
      "fixed_ip": null,
      "id": "1",
      "instance_id": null,
      "ip": "10.10.10.1",
      "pool": "nova"
    },
    {
      "fixed_ip": null,
      "id": "2",
      "instance_id": null,
      "ip": "10.10.10.2",
      "pool": "nova"
    }
  ]
}
`

// FirstFloatingIP is the first result in ListOutput.
var FirstFloatingIP = FloatingIP{
	ID:         "1",
	IP:         "10.10.10.1",
	InstanceID: "",
	Pool:       "nova",
	FixedIP:    "",
}

// SecondFloatingIP is the second result in ListOutput.
var SecondFloatingIP = FloatingIP{
	ID:         "2",
	IP:         "10.10.10.2",
	InstanceID: "",
	Pool:       "nova",
	FixedIP:    "",
}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-floating-ips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
