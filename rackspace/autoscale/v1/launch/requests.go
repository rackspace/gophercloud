package launch

import "github.com/rackspace/gophercloud"

// Get requests the details of a given auto scale group's launch configuration.
func Get(client *gophercloud.ServiceClient, groupID string) GetResult {
	var result GetResult

	_, result.Err = client.Get(launchURL(client, groupID), &result.Body, nil)

	return result
}
