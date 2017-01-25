package projects

import "github.com/rackspace/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

func pairProjectGroupAndRoleURL(client *gophercloud.ServiceClient, dID, gID, rID string) string {
	return client.ServiceURL("projects/" + dID + "/groups/" + gID + "/roles/" + rID)
}
