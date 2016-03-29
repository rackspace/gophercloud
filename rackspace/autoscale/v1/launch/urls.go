package launch

import "github.com/rackspace/gophercloud"

func launchURL(c *gophercloud.ServiceClient, groupID string) string {
	return c.ServiceURL("groups", groupID, "launch")
}
