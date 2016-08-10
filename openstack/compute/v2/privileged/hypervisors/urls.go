package hypervisors

import (
	"github.com/rackspace/gophercloud"
	"fmt"
)


func getListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-hypervisors")
}

func getDetailedListURL(client *gophercloud.ServiceClient) string{
	return client.ServiceURL("os-hypervisors/detail")
}

func getDetailURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(fmt.Sprintf("os-hypervisors/%s", id))
}

func getHypervisorServersURL(client *gophercloud.ServiceClient, hypervisorHostname string) string {
	return client.ServiceURL(fmt.Sprintf("os-hypervisors/%s/servers", hypervisorHostname))
}

func getHypervisorUptimeURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(fmt.Sprintf("os-hypervisors/%s/uptime", id))
}
