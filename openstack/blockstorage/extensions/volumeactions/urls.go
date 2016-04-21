package volumeactions

import "github.com/rackspace/gophercloud"

func volumeActionsURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
