package hypervisors

import (
	"github.com/rackspace/gophercloud"
)

// TODO: introduce list options for the pagination - marker, limit, page_size
// as this is defined in gophercloud/openstack/v2/servers/requests.go:{row 25}
// currently OpenStack paging only supports servers and flavors

// List makes a request against the API to list hypervisors accessible to you
// in pagination way (Not supported now by OpenStack for hypervisors).
func List(client *gophercloud.ServiceClient) GetResult {
	var result GetResult
	_, result.Err = client.Get(getListURL(client),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}

// Get details about specified by id hypervisor
func GetDetailsList(client *gophercloud.ServiceClient) GetResult {
	var result GetResult
	_, result.Err = client.Get(getDetailedListURL(client), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return result
}

// Get details about specified by id hypervisor
func GetDetail(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getDetailURL(client, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return result
}

// Get VMs booted on specified by hostname hypervisor
func GetHypervisorServers(client *gophercloud.ServiceClient, hypervisorHostname string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getHypervisorServersURL(client, hypervisorHostname),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}

func GetHypervisorUptime(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getHypervisorUptimeURL(client, id),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}
