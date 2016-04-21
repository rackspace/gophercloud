package volumeactions

import (
	fixtures "github.com/rackspace/gophercloud/openstack/blockstorage/extensions/volumeactions/testing"
	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
	"testing"
)

func TestReserve(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixtures.MockReserveResponse(t)
	res := Reserve(client.ServiceClient(), "58003305-1778-43ce-ac78-a81fe255db15")
	th.AssertNoErr(t, res.Err)
}

func TestInitializeConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixtures.MockInitializeConnectionResponse(t)
	connector := &ConnectorOpts{
		IP:        "192.168.0.37",
		Host:      "devbox",
		Initiator: "iqn.1993-08.org.debian:01:17a0e6ac38f8",
		MultiPath: false,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	res := InitializeConnection(client.ServiceClient(), "58003305-1778-43ce-ac78-a81fe255db15", connector)
	th.AssertNoErr(t, res.Err)
}

func TestAttach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixtures.MockAttachResponse(t)
	options := &AttachOpts{
		MountPoint:   "/dev/vdb",
		Mode:         "rw",
		InstanceUUID: "4e6b240c-e32a-4e7b-8453-e4ed0f5eb107",
	}
	res := Attach(client.ServiceClient(), "58003305-1778-43ce-ac78-a81fe255db15", options)
	th.AssertNoErr(t, res.Err)
}

func TestUnReserve(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixtures.MockUnReserveResponse(t)
	res := UnReserve(client.ServiceClient(), "58003305-1778-43ce-ac78-a81fe255db15")
	th.AssertNoErr(t, res.Err)
}

func TestTermianteConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixtures.MockTerminateConnectionResponse(t)
	connector := &ConnectorOpts{
		IP:        "192.168.0.37",
		Host:      "devbox",
		Initiator: "iqn.1993-08.org.debian:01:17a0e6ac38f8",
		MultiPath: false,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	res := TerminateConnection(client.ServiceClient(), "58003305-1778-43ce-ac78-a81fe255db15", connector)
	th.AssertNoErr(t, res.Err)
}

func TestDetach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixtures.MockUnReserveResponse(t)
	res := UnReserve(client.ServiceClient(), "58003305-1778-43ce-ac78-a81fe255db15")
	th.AssertNoErr(t, res.Err)
}
