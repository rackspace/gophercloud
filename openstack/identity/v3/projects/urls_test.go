package projects

import (
	"testing"

	"github.com/rackspace/gophercloud"
)

func TestGetURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := getURL(&client, "12345")
	if url != "http://localhost:5000/v3/projects/12345" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}

func TestListURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := listURL(&client)
	if url != "http://localhost:5000/v3/projects" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}

func TestCreateURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := createURL(&client)
	if url != "http://localhost:5000/v3/projects" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}

func TestUpdateURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := updateURL(&client, "12345")
	if url != "http://localhost:5000/v3/projects/12345" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}

func TestDeleteURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := deleteURL(&client, "12345")
	if url != "http://localhost:5000/v3/projects/12345" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}
