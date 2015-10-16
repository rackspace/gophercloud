package query

import "github.com/rackspace/gophercloud"

var root = "views"

func getRootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(root)
}

func getURLForListAndPoints(c *gophercloud.ServiceClient) string {
	return getRootURL(c)
}

func getURLForListAndResolution(c *gophercloud.ServiceClient) string {
	return getRootURL(c)
}

func getURLForPoints(c *gophercloud.ServiceClient, metric string) string {
	return c.ServiceURL(root, metric)
}

func getURLForResolution(c *gophercloud.ServiceClient, metric string) string {
	return c.ServiceURL(root, metric)
}

func getEventURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events", "getEvents")
}

func getSearchURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics", "search")
}

func getLimits(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("limits")
}

