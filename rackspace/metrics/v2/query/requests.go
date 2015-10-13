package query

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
)

const (
	SUM = "sum"
	MAX = "max"
	MIN = "min"
	LATEST = "latest"
	VARIANCE = "variance"
	AVERAGE = "averge"
)

//Roll-ups process full-resolution data into coarser granularities of 5 minutes, 20 minutes, 60 minutes, 4 hours and 24 hours.
const (
	FULL = "full"
	MIN5 = "min5"
	MIN20 = "min20"
	MIN60 = "min60"
	MIN240 = "min240"
	MIN1440 = "min1440"
)

// GetDataForListByPoints retrieve data against a list of metrics and number of points, for the specified tenant associated with RackspaceMetrics.
func GetDataForListByPoints(c *gophercloud.ServiceClient, from int64, to int64, points int32, metrics ...string) (MetricListData, error) {
	var res GetResult
	reqBody := make([]interface{}, len(metrics))
	for i, v := range metrics {
		reqBody[i] = v
	}
	_, res.Err = c.Post(getURLForListAndPoints(c, from, to, points), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	b := res.Body.(interface{})
	var metricList MetricListData
	err := mapstructure.Decode(b, &metricList)
	return metricList, err
}

// GetDataForListByResolution retrieve data against a list of metrics and specified resolution, for the specified tenant associated with RackspaceMetrics.
func GetDataForListByResolution(c *gophercloud.ServiceClient, from int64, to int64, resolution string, metrics ...string) (MetricListData, error) {
	var res GetResult
	reqBody := make([]interface{}, len(metrics))
	for i, v := range metrics {
		reqBody[i] = v
	}
	_, res.Err = c.Post(getURLForListAndResolution(c, from, to, resolution), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	b := res.Body.(interface{})
	var metricList MetricListData
	err := mapstructure.Decode(b, &metricList)
	return metricList, err
}

// GetDataByPoints retrieve metric data by points, for the specified tenant associated with RackspaceMetrics.
func GetDataByPoints(c *gophercloud.ServiceClient, metric string, from int64, to int64, points int32, sel ...string) GetResult {
	var res GetResult
	if len(sel) >= 1 {
		_, res.Err = c.Get(getURLForPointsWithSelect(c, metric, from, to, points, sel), &res.Body, nil)
		return res
	}
	_, res.Err = c.Get(getURLForPoints(c, metric, from, to, points), &res.Body, nil)
	return res
}

// GetDataByPoints retrieve metric data by resolution, for the specified tenant associated with RackspaceMetrics.
func GetDataByResolution(c *gophercloud.ServiceClient, metric string, from int64, to int64, resolution string, sel ...string) GetResult {
	var res GetResult
	if len(sel) >= 1 {
		_, res.Err = c.Get(getURLForResolutionWithSelect(c, metric, from, to, resolution, sel), &res.Body, nil)
		return res
	}
	_, res.Err = c.Get(getURLForResolution(c, metric, from, to, resolution), &res.Body, nil)
	return res
}

// SearchMetric retrieves a list of available metrics for the specified tenant associated with RackspaceMetrics.
func SearchMetric(c *gophercloud.ServiceClient, metric string) ([]Metric, error) {
	var res GetResult
	_, res.Err = c.Get(getSearchURL(c, metric), &res.Body, nil)
	b := res.Body.(interface{})
	var metrics []Metric
	err := mapstructure.Decode(b, &metrics)
	return metrics, err
}

//GetEvents retrieves a list of events for the specified tenant associated with RackspaceMetrics.
func GetEvents(c *gophercloud.ServiceClient, from int64, until int64, tag ...string) ([]Event, error) {
	var res GetResult
	if len(tag) == 1 {
		_, res.Err = c.Get(getEventURLForTag(c, from, until, tag[0]), &res.Body, nil)
	} else {
		_, res.Err = c.Get(getEventURL(c, from, until), &res.Body, nil)
	}
	b := res.Body.(interface{})
	var events []Event
	err := mapstructure.Decode(b, &events)
	return events, err
}

// GetLimits retrieves the number of API transaction that are available for the specified tenant associated with RackspaceMetrics.
func GetLimits(c *gophercloud.ServiceClient) (Limits, error) {
	var res GetResult
	_, res.Err = c.Get(getLimits(c), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 404},
	})
	b := res.Body.(map[string]interface{})
	var limits Limits
	err := mapstructure.Decode(b["limits"], &limits)
	return limits, err
}