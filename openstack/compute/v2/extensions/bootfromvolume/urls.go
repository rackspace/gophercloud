package bootfromvolume

import "github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
