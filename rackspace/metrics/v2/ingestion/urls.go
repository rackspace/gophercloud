package ingestion

import "github.com/rackspace/gophercloud"

var root = "ingest"

func getURLForIngestMetrics(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(root)
}

func getURLForIngestAggregatedMetrics(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(root, "aggregated")
}

func getURLForIngestEvents(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}