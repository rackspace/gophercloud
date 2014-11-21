package floatingips

import "github.com/rackspace/gophercloud"

const resourcePath = "os-floating-ips"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func allocateURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func deallocateURL(c *gophercloud.ServiceClient, id string) string {
	return getURL(c, id)
}

func addURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, "action")
}

func removeURL(c *gophercloud.ServiceClient, serverID string) string {
	return addURL(c, serverID)
}
