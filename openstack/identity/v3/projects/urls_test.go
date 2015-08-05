package projects

import (
	"testing"

	"github.com/rackspace/gophercloud"
)

func TestListURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := listURL(&client)
	if url != "http://localhost:5000/v3/projects" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}

func TestProjectURL(t *testing.T) {
	client := gophercloud.ServiceClient{Endpoint: "http://localhost:5000/v3/"}
	url := projectURL(&client, "12345")
	if url != "http://localhost:5000/v3/projects/12345" {
		t.Errorf("Unexpected project URL generated: [%s]", url)
	}
}
