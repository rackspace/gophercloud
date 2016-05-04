package launch

import "github.com/rackspace/gophercloud"

func getURL(c *gophercloud.ServiceClient, groupID string) string {
	return c.ServiceURL("groups", groupID, "launch")
}

func updateURL(c *gophercloud.ServiceClient, groupID string) string {
	return getURL(c, groupID)
}
