package images

import (
	"fmt"
	"github.com/rackspace/gophercloud"
)

func addURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images")
}

func uploadUrl(client *gophercloud.ServiceClient, imageId string) string {
	return client.ServiceURL(fmt.Sprintf("images/%s/file", imageId))
}
