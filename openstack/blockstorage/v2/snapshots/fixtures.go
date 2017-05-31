package snapshots

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "snapshot": {
        "name": "snap-001",
        "volume_id": "32d8295e-17ef-4ea6-9179-eb71f6827f20",
        "description": "test-snapshot",
		"force": false
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "snapshot": {
        "name": "snap-001",
        "volume_id": "32d8295e-17ef-4ea6-9179-eb71f6827f20",
        "desription": "test-snapshot",
		"size": 1,
		"id": "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af"
    }
}
	`)
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/4ee8a3f6-d1c8-4541-ad09-06b7e84a68af", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
  "snapshot":
    {"status": "available",
     "description": "test-snapshot",
     "updated_at": "2017-05-31T14:18:36.000000",
     "volume_id": "32d8295e-17ef-4ea6-9179-eb71f6827f20",
     "id": "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af",
     "size": 1,
     "name": "snap-001",
     "created_at": "2017-05-31T14:18:35.000000",
     "metadata": {}
    }
}
      `)
	})
}

func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/4ee8a3f6-d1c8-4541-ad09-06b7e84a68af", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
  {
  "snapshots": [
    {
      "status": "available",
      "description": null,
      "updated_at": "2017-05-31T14:18:36.000000",
      "volume_id": "32d8295e-17ef-4ea6-9179-eb71f6827f20",
      "id": "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af",
      "size": 1,
      "name": null,
      "created_at": "2017-05-31T14:18:35.000000",
      "metadata": {"foo": "bar"}
    },
    {
      "status": "available",
      "description": "this is only a test snapshot",
      "updated_at": "2017-05-31T14:10:13.000000",
      "volume_id": "32d8295e-17ef-4ea6-9179-eb71f6827f20",
      "id": "c970ff21-3c2b-4a4c-b6a0-731808a81776",
      "size": 1,
      "name": "test-snap",
      "created_at": "2017-05-31T14:10:12.000000",
      "metadata": {}
      }
  ]
  }
  `)
	})
}
