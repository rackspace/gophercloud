package launch

import (
	"testing"

	"github.com/rackspace/gophercloud"
	th "github.com/rackspace/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient(), "123")
	expected := endpoint + "groups/123/launch"
	th.CheckEquals(t, expected, actual)
}

func TestUpdateURL(t *testing.T) {
	actual := updateURL(endpointClient(), "123")
	expected := endpoint + "groups/123/launch"
	th.CheckEquals(t, expected, actual)
}
