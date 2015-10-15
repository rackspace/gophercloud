package ingestion

import "github.com/rackspace/gophercloud"

// SendMetrics will ingest the metrics for the tenant associated with RackspaceMetrics.
func SendMetrics(c *gophercloud.ServiceClient, metrics []MetricData) {
	var res PostResult
	reqBody := make([]map[string]interface{}, len(metrics))
	for i := range metrics {
		Metric := metrics[i]
		reqBody[i] = map[string]interface{}{
			"collectionTime": Metric.CollectionTime,
			"ttlInSeconds":Metric.TtlInSeconds,
			"metricValue": Metric.MetricValue,
			"metricName": Metric.MetricName,
		}
	}
	_, res.Err = c.Post(getURLForIngestMetrics(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}

// SendMetrics will ingest the pre aggregated metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedMetrics(c *gophercloud.ServiceClient, metrics AggregatedMetricData) {
	var res PostResult

	reqBody := map[string]interface{}{
		"tenantId": metrics.TenantId,
		"timestamp": metrics.Timestamp,
		"counters": convertCounters(metrics.Counters),
		"timers": convertTimers(metrics.Timers),
		"gauges": convertGauges(metrics.Gauges),
		"sets": convertSets(metrics.Sets),
	}

	_, res.Err = c.Post(getURLForIngestAggregatedMetrics(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}

// SendMetrics will ingest the pre aggregated counter metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedCounters(c *gophercloud.ServiceClient, tenantId string, timestamp int64, Counters []Counter) {
	var res PostResult

	counters := convertCounters(Counters)
	reqBody := map[string]interface{}{
		"tenantId": tenantId,
		"timestamp": timestamp,
		"counters": counters,
	}

	_, res.Err = c.Post(getURLForIngestAggregatedMetrics(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}

// SendMetrics will ingest the pre aggregated timer metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedTimers(c *gophercloud.ServiceClient, tenantId string, timestamp int64, Timers []Timer) {
	var res PostResult

	timers := convertTimers(Timers)
	reqBody := map[string]interface{}{
		"tenantId": tenantId,
		"timestamp": timestamp,
		"timers": timers,
	}
	_, res.Err = c.Post(getURLForIngestAggregatedMetrics(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}

// SendMetrics will ingest the pre aggregated gauge metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedGauges(c *gophercloud.ServiceClient, tenantId string, timestamp int64, Gauges []Gauge) {
	var res PostResult

	gauges := convertGauges(Gauges)
	reqBody := map[string]interface{}{
		"tenantId": tenantId,
		"timestamp": timestamp,
		"gauges": gauges,
	}

	_, res.Err = c.Post(getURLForIngestAggregatedMetrics(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}

// SendMetrics will ingest the pre aggregated set metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedSets(c *gophercloud.ServiceClient, tenantId string, timestamp int64, Sets []Set) {
	var res PostResult

	sets := convertSets(Sets)
	reqBody := map[string]interface{}{
		"tenantId": tenantId,
		"timestamp": timestamp,
		"sets": sets,
	}

	_, res.Err = c.Post(getURLForIngestAggregatedMetrics(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}

// SendMetrics will ingest the special events for the tenant associated with RackspaceMetrics.
func SendEvent(c *gophercloud.ServiceClient, event Event) {
	var res PostResult

	reqBody := map[string]interface{}{
		"what": event.What,
		"when": event.When,
		"tags": event.Tags,
		"data": event.Data,
	}

	_, res.Err = c.Post(getURLForIngestEvents(c), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
}