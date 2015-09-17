package ipaddresses

import "github.com/rackspace/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("ip_addresses", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ip_addresses")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listByServerURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, "ip_associations")
}

func getByServerURL(c *gophercloud.ServiceClient, serverID, sharedIPID string) string {
	return c.ServiceURL("servers", serverID, "ip_associations", sharedIPID)
}

func associateURL(c *gophercloud.ServiceClient, serverID, sharedIPID string) string {
	return getByServerURL(c, serverID, sharedIPID)
}

func disassociateURL(c *gophercloud.ServiceClient, serverID, sharedIPID string) string {
	return getByServerURL(c, serverID, sharedIPID)
}
