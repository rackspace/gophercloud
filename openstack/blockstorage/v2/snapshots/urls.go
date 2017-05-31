package snapshots

import "github.com/rackspace/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots", "detail")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}
