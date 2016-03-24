package groups

import "github.com/rackspace/gophercloud"

const (
	groups = "groups"
)

func groupsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(groups)
}
