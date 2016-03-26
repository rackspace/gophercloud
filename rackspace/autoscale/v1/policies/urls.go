package policies

import "github.com/rackspace/gophercloud"

func policiesURL(c *gophercloud.ServiceClient, groupID string) string {
	return c.ServiceURL("groups", groupID, "policies")
}
