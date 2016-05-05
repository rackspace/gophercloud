package groups

import "github.com/rackspace/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("groups")
}

func stateURL(c *gophercloud.ServiceClient, groupID string) string {
	return c.ServiceURL("groups", groupID, "state")
}

func configURL(c *gophercloud.ServiceClient, groupID string) string {
	return c.ServiceURL("groups", groupID, "config")
}
