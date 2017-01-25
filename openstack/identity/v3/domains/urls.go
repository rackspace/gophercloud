package domains

import "github.com/rackspace/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("domains")
}

func pairDomainGroupAndRoleURL(client *gophercloud.ServiceClient, dID, gID, rID string) string {
	return client.ServiceURL("domains/" + dID + "/groups/" + gID + "/roles/" + rID)
}
