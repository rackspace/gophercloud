package images

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
	actual := getURL(endpointClient(), "foo")
	expected := endpoint + "images/foo"
	th.CheckEquals(t, expected, actual)
}

func TestListDetailURL(t *testing.T) {
	actual := listDetailURL(endpointClient())
	expected := endpoint + "images/detail"
	th.CheckEquals(t, expected, actual)
}

func TestMetadataURL(t *testing.T) {
	actual := metadataURL(endpointClient(), "foo")
	expected := endpoint + "images/foo/metadata"
	th.CheckEquals(t, expected, actual)
}

func TestMetadatumURL(t *testing.T) {
	actual := metadatumURL(endpointClient(), "foo", "bar")
	expected := endpoint + "images/foo/metadata/bar"
	th.CheckEquals(t, expected, actual)
}
