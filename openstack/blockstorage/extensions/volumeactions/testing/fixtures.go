package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func MockAttachResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-attach":
    {
        "mountpoint": "/dev/vdb",
        "mode": "rw",
        "instance_uuid": "4e6b240c-e32a-4e7b-8453-e4ed0f5eb107"
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockReserveResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-reserve": {}
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockUnreserveResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-unreserve": {}
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockInitializeConnectionResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-initialize_connection":
    {
        "connector":
        {
        "ip":"192.168.0.37",
        "host":"devbox",
		"initiator":"iqn.1993-08.org.debian:01:17a0e6ac38f8",
        "multipath": false,
        "platform": "x86_64",
        "os_type": "linux2"
        }
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `
{
    "connection_info":
    {
        "driver_volume_type": "iscsi",
        "data":
        {
            "target_discovered": false,
            "encrypted": false,
            "target_iqn": "iqn.2010-10.org.openstack:volume-58003305-1778-43ce-ac78-a81fe255db15",
            "target_portal": "192.168.0.37:3260",
            "volume_id": "58003305-1778-43ce-ac78-a81fe255db15",
            "target_lun": 1,
            "access_mode": "rw",
            "auth_username": "fRYNAWbGR7HRYwUhHGa6",
            "auth_password": "J3CEGp2EEtfAqk4T",
            "auth_method": "CHAP",
			"qos_specs": null
        }
    }
}`)
		})
}

func MockTerminateConnectionResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
	"os-terminate_connection":
	{
		"connector":
		{
			"initiator": "iqn.1993-08.org.debian:01:17a0e6ac38f8",
			"ip": "192.168.0.37",
			"platform": "x86_64",
			"host": "devbox",
			"os_type": "linux2",
			"multipath": false
		}
	}
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockUnReserveResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-unreserve": {}
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockDetachResponse(t *testing.T) {
	// TOD(jdg): Add optional attach_id, for now it's not interesting because
	// Cinder doesn't actually provide it; BUT future proposal is to return it
	// to caller on initialize_connection
	th.Mux.HandleFunc("/volumes/58003305-1778-43ce-ac78-a81fe255db15/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-detach": {}
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}
