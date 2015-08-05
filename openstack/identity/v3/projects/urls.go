package projects

import "github.com/rackspace/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("projects")
}

func projectURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}
