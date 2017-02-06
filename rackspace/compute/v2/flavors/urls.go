package flavors

import (
	"github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"
)

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}
