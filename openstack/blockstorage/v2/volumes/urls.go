package volumes

import "github.com/rackspace/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes", "detail")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func attachURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}

func detachURL(c *gophercloud.ServiceClient, id string) string {
	return attachURL(c, id)
}

func reserveURL(c *gophercloud.ServiceClient, id string) string {
	return attachURL(c, id)
}

func unreserveURL(c *gophercloud.ServiceClient, id string) string {
	return attachURL(c, id)
}

func initializeConnectionURL(c *gophercloud.ServiceClient, id string) string {
	return attachURL(c, id)
}

func teminateConnectionURL(c *gophercloud.ServiceClient, id string) string {
	return attachURL(c, id)
}
