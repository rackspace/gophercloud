package images

import "github.com/rackspace/gophercloud"

func listDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images", "detail")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func metadataURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id, "metadata")
}

func metadatumURL(client *gophercloud.ServiceClient, id, key string) string {
	return client.ServiceURL("images", id, "metadata", key)
}
