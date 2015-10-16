package ingestion

import (
	"github.com/rackspace/gophercloud"
	"fmt"
)

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

func convertCounters(Counters []Counter) []map[string]interface{} {

	counters := make([]map[string]interface{}, len(Counters))
	for i := range Counters {
		Counter := Counters[i]
		counters[i] = map[string]interface{}{
			"name": Counter.Name,
			"rate": Counter.Rate,
			"value": Counter.Value,
		}
	}

	return counters
}

func convertTimers(Timers []Timer) []map[string]interface{} {

	timers := make([]map[string]interface{}, len(Timers))
	for i := range Timers {
		Timer := Timers[i]

		Percentiles := Timer.Percentiles
		percentiles := make(map[string]interface{}, len(Percentiles))
		for j := range Percentiles {
			Percentile := Percentiles[j]
			percentiles[fmt.Sprintf("%d", Percentile.Key)] =
			map[string]interface{}{
				"avg": Percentile.Value.Average,
				"max": Percentile.Value.Max,
				"sum": Percentile.Value.Sum,
			}
		}

		Histograms := Timer.Histograms
		histograms := make(map[string]interface{}, len(Histograms))
		for j := range Histograms {
			Histogram := Histograms[j]
			histograms["bin_"+Histogram.Bin] = Histogram.Frequency
		}

		timers[i] = map[string]interface{}{
			"name": Timer.Name,
			"count": Timer.Count,
			"rate": Timer.Rate,
			"min": Timer.Min,
			"max": Timer.Max,
			"sum": Timer.Sum,
			"avg": Timer.Average,
			"median": Timer.Median,
			"std": Timer.Std,
			"percentiles": percentiles,
			"histogram": histograms,
		}
	}

	return timers
}

func convertGauges(Gauges []Gauge) []map[string]interface{} {

	gauges := make([]map[string]interface{}, len(Gauges))
	for i := range Gauges {
		Gauge := Gauges[i]
		gauges[i] = map[string]interface{}{
			"name": Gauge.Name,
			"value": Gauge.Value,
		}
	}
	return gauges
}

func convertSets(Sets []Set) []map[string]interface{} {

	sets := make([]map[string]interface{}, len(Sets))
	for i := range Sets {
		Set := Sets[i]
		sets[i] = map[string]interface{}{
			"name": Set.Name,
			"values": Set.Values,
		}
	}
	return sets
}
