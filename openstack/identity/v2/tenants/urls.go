package tenants

import "github.com/rackspace/gophercloud"

func ResourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("tenants", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("tenants")
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
