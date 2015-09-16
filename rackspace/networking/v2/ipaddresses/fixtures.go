package sharedips

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/ip_addresses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
			{
		    "ip_addresses": [
		       {
		            "id": "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
		            "network_id": "6870304a-7212-443f-bd0c-089c886b44df",
		            "address": "192.168.10.1",
		            "port_ids": [
									"2f693cca-7383-45da-8bae-d26b6c2d6718"
								],
		            "subnet_id": "f11687e8-ef0d-4207-8e22-c60e737e473b",
		            "tenant_id": "2345678",
		            "version": "4",
		            "type": "fixed"
		        }
		    ]
		}
  `)
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/ip_addresses/4cacd68e-d7aa-4ff2-96f4-5c6f57dba737", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
			{
			    "ip_address":
			    {
			        "id": "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
			        "network_id": "fda61e0b-a410-49e8-ad3a-64c595618c7e",
			        "address": "192.168.10.1",
			        "port_ids": ["6200d533-a42b-4c04-82a1-cc14dbdbf2de",
			                    "9d0db2d7-62df-4c99-80cb-6f140a5260e8"],
			        "subnet_id": "f11687e8-ef0d-4207-8e22-c60e737e473b",
			        "tenant_id": "2345678",
			        "version": "4",
			        "type": "shared"
			    }
			}
    `)
	})
}

func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc("/ip_addresses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
			{
		    "ip_address": {
		        "network_id": "00000000-0000-0000-0000-000000000000",
		        "version": 4,
		        "port_ids": [
		            "6200d533-a42b-4c04-82a1-cc14dbdbf2de",
		            "9d0db2d7-62df-4c99-80cb-6f140a5260e8"
		         ],
		        "tenant_id": "2345678"
		    }
			}
    `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, `
			{
		    "ip_address":
		    {
		        "id": "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
		        "network_id": "fda61e0b-a410-49e8-ad3a-64c595618c7e",
		        "address": "192.168.10.1",
		        "port_ids": ["6200d533-a42b-4c04-82a1-cc14dbdbf2de",
		                    "9d0db2d7-62df-4c99-80cb-6f140a5260e8"],
		        "subnet_id": "f11687e8-ef0d-4207-8e22-c60e737e473b",
		        "tenant_id": "2345678",
		        "version": "4",
		        "type": "shared"
		    }
			}
    `)
	})
}

func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc("/ip_addresses/4cacd68e-d7aa-4ff2-96f4-5c6f57dba737", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}

func MockUpdateResponse(t *testing.T) {
	th.Mux.HandleFunc("/ip_addresses/4cacd68e-d7aa-4ff2-96f4-5c6f57dba737", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, `
			{
    "ip_address": {
        "port_ids": ["275b0516-206f-4421-8e42-1d3d1e4e9fb2", "66811c0a-fdbd-49d4-b1dd-f0f15a329744"]
    }
}
		`)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
			{
			    "ip_address":
			    {
			        "id": "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
			        "network_id": "6870304a-7212-443f-bd0c-089c886b44df",
			        "address": "192.168.10.1",
			        "port_ids": ["275b0516-206f-4421-8e42-1d3d1e4e9fb2",
			                    "66811c0a-fdbd-49d4-b1dd-f0f15a329744"],
			        "subnet_id": "f11687e8-ef0d-4207-8e22-c60e737e473b",
			        "tenant_id": "2345678",
			        "version": "4",
			        "type": "shared"
			    }
			}
    `)
	})
}

func MockListByServerResponse(t *testing.T) {
	th.Mux.HandleFunc("/servers/123456/ip_associations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
			{
			    "ip_associations":
			    [
			        {
			            "id": "1",
			            "address": "10.1.1.1"
			        },
			        {
			            "id": "2",
			            "address": "10.1.1.2"
			        }
			    ]
			}
  `)
	})
}

func MockGetByServerResponse(t *testing.T) {
	th.Mux.HandleFunc("/servers/123456/ip_associations/1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
			{
    "ip_association":
    {
        "id": "1",
        "address": "10.1.1.1"
    }
}
      `)
	})
}

func MockAssociateResponse(t *testing.T) {
	th.Mux.HandleFunc("/servers/123456/ip_associations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, `
			{
			    "ip_association":
			        {
			            "id": "2",
			            "address": "10.1.1.2"
			        }
			}
    `)
	})
}

func MockDisassociateResponse(t *testing.T) {
	th.Mux.HandleFunc("/servers/123456/ip_associations/2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
