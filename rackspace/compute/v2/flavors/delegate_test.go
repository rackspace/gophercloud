package flavors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestListFlavors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, ListOutput)
		case "performance1-2":
			fmt.Fprintf(w, `{ "flavors": [] }`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})

	count := 0
	err := List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		actual, err := ExtractFlavors(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedFlavorSlice, actual)

		count++
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
