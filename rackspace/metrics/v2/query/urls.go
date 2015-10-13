package query

import (
	"github.com/rackspace/gophercloud"
	"fmt"
	"strings"
)

var root = "views"

func getURLForListAndPoints(c *gophercloud.ServiceClient, from int64, to int64, points int32) string {
	return c.ServiceURL(root+"?from="+fmt.Sprintf("%d", from)+"&to="+fmt.Sprintf("%d", to)+"&points="+fmt.Sprintf("%d", points))
}

func getURLForListAndResolution(c *gophercloud.ServiceClient, from int64, to int64, resolution string) string {
	return c.ServiceURL(root+"?from="+fmt.Sprintf("%d", from)+"&to="+fmt.Sprintf("%d", to)+"&resolution="+resolution)
}

func getURLForPoints(c *gophercloud.ServiceClient, metric string, from int64, to int64, points int32) string {
	return c.ServiceURL(root, metric+"?from="+fmt.Sprintf("%d", from)+"&to="+fmt.Sprintf("%d", to)+"&points="+fmt.Sprintf("%d", points))
}

func getURLForResolution(c *gophercloud.ServiceClient, metric string, from int64, to int64, resolution string) string {
	return c.ServiceURL(root, metric+"?from="+fmt.Sprintf("%d", from)+"&to="+fmt.Sprintf("%d", to)+"&resolution="+resolution)
}

func getURLForPointsWithSelect(c *gophercloud.ServiceClient, metric string, from int64, to int64, points int32, sel []string) string {
	return c.ServiceURL(root, metric+"?from="+fmt.Sprintf("%d", from)+"&to="+fmt.Sprintf("%d", to)+"&points="+fmt.Sprintf("%d", points)+"&select="+strings.Join(sel, "&select="))
}

func getURLForResolutionWithSelect(c *gophercloud.ServiceClient, metric string, from int64, to int64, resolution string, sel []string) string {
	return c.ServiceURL(root, metric+"?from="+fmt.Sprintf("%d", from)+"&to="+fmt.Sprintf("%d", to)+"&resolution="+resolution+"&select="+strings.Join(sel, "&select="))
}

func getEventURL(c *gophercloud.ServiceClient, from int64, until int64) string {
	return c.ServiceURL("events", "getEvents?from="+fmt.Sprintf("%d", from)+"&until="+fmt.Sprintf("%d", until))
}

func getEventURLForTag(c *gophercloud.ServiceClient, from int64, until int64, tag string) string {
	return c.ServiceURL("events", "getEvents?from="+fmt.Sprintf("%d", from)+"&until="+fmt.Sprintf("%d", until)+"&tags="+tag)
}

func getSearchURL(c *gophercloud.ServiceClient, metric string) string {
	return c.ServiceURL("metrics", "search?query="+metric)
}

func getLimits(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("limits")
}

