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

// SendAggregatedMetrics will ingest the pre aggregated metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedMetrics(c *gophercloud.ServiceClient, metrics AggregatedMetricData) {
	SendAggregatedMetricsData(c, metrics)
}

// SendAggregatedCounters will ingest the pre aggregated counter metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedCounters(c *gophercloud.ServiceClient, tenantId string, timestamp int64, counters []Counter) {

	aggregatedMetaData := AggregatedMetricData{
		TenantId:tenantId,
		Timestamp:timestamp,
		Counters:counters,
	}

	SendAggregatedMetricsData(c, aggregatedMetaData)
}

// SendAggregatedTimers will ingest the pre aggregated timer metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedTimers(c *gophercloud.ServiceClient, tenantId string, timestamp int64, timers []Timer) {
	aggregatedMetaData := AggregatedMetricData{
		TenantId:tenantId,
		Timestamp:timestamp,
		Timers:timers,
	}

	SendAggregatedMetricsData(c, aggregatedMetaData)
}

// SendAggregatedGauges will ingest the pre aggregated gauge metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedGauges(c *gophercloud.ServiceClient, tenantId string, timestamp int64, gauges []Gauge) {

	aggregatedMetaData := AggregatedMetricData{
		TenantId:tenantId,
		Timestamp:timestamp,
		Gauges:gauges,
	}

	SendAggregatedMetricsData(c, aggregatedMetaData)
}

// SendAggregatedSets will ingest the pre aggregated set metrics for the tenant associated with RackspaceMetrics.
func SendAggregatedSets(c *gophercloud.ServiceClient, tenantId string, timestamp int64, sets []Set) {

	aggregatedMetaData := AggregatedMetricData{
		TenantId:tenantId,
		Timestamp:timestamp,
		Sets:sets,
	}

	SendAggregatedMetricsData(c, aggregatedMetaData)
}

// SendEvent will ingest the special events for the tenant associated with RackspaceMetrics.
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

//To collectively send aggregated metrics.
func SendAggregatedMetricsData(c *gophercloud.ServiceClient, aggregatedMetricData AggregatedMetricData) {
	var res PostResult

	aggregatedMetricDataMap := aggregatedMetricData.ToAggregatedMetricDataMap()

	_, res.Err = c.Post(getURLForIngestAggregatedMetrics(c), aggregatedMetricDataMap, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

}

//Transforms AggregatedMetricData (Struct) to Map.
func (aggregatedMetricData AggregatedMetricData) ToAggregatedMetricDataMap() (map[string]interface{}) {
	aggregatedMetricDataMap := make(map[string]interface{})

	aggregatedMetricDataMap["tenantId"] = aggregatedMetricData.TenantId
	aggregatedMetricDataMap["timestamp"] = aggregatedMetricData.Timestamp

	if aggregatedMetricData.Counters != nil {
		Counters := aggregatedMetricData.Counters
		counters := make([]map[string]interface{}, len(Counters))
		for i := range Counters {
			Counter := Counters[i]
			counters[i] = map[string]interface{}{
				"name": Counter.Name,
				"rate": Counter.Rate,
				"value": Counter.Value,
			}
		}
		aggregatedMetricDataMap["counters"] = counters
	}

	if aggregatedMetricData.Timers != nil {
		Timers := aggregatedMetricData.Timers
		timers := make([]map[string]interface{}, len(Timers))
		for i := range Timers {
			Timer := Timers[i]
			timer := make(map[string]interface{})

			timer["name"] = Timer.Name
			timer["count"] = Timer.Count
			timer["rate"] = Timer.Rate
			timer["min"] = Timer.Min
			timer["max"] = Timer.Max
			timer["sum"] = Timer.Sum
			timer["avg"] = Timer.Average
			timer["median"] = Timer.Median
			timer["std"] = Timer.Std

			if Timer.Histograms != nil {

				Histograms := Timer.Histograms
				histograms := make(map[string]interface{}, len(Histograms))
				for j := range Histograms {
					Histogram := Histograms[j]
					histograms["bin_"+Histogram.Bin] = Histogram.Frequency
				}
				timer["histogram"] = histograms
			}

			if Timer.Percentiles != nil {

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
				timer["percentiles"] = percentiles
			}

			timers[i] = timer
		}
		aggregatedMetricDataMap["timers"] = timers
	}

	if aggregatedMetricData.Gauges != nil {
		Gauges := aggregatedMetricData.Gauges
		gauges := make([]map[string]interface{}, len(Gauges))
		for i := range Gauges {
			Gauge := Gauges[i]
			gauges[i] = map[string]interface{}{
				"name": Gauge.Name,
				"value": Gauge.Value,
			}
		}
		aggregatedMetricDataMap["gauges"] = gauges
	}

	if aggregatedMetricData.Sets != nil {
		Sets := aggregatedMetricData.Sets
		sets := make([]map[string]interface{}, len(Sets))
		for i := range Sets {
			Set := Sets[i]
			sets[i] = map[string]interface{}{
				"name": Set.Name,
				"values": Set.Values,
			}
		}
		aggregatedMetricDataMap["sets"] = sets
	}
	return aggregatedMetricDataMap
}