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

type QueryParams struct {
	From       int64 `q:"from"`
	To         int64 `q:"to"`
	Until      int64 `q:"until"`
	Points     int32  `q:"points"`
	Resolution string `q:"resolution"`
	Select     []string `q:"select"`
	Query      string `q:"query"`
	Tags       string `q:"tags"`
}


// GetDataForListByPoints retrieve data against a list of metrics and number of points, for the specified tenant associated with RackspaceMetrics.
func GetDataForListByPoints(c *gophercloud.ServiceClient, opts QueryParams, metrics ...string) (MetricListData, error) {
	var res GetResult

	reqBody := make([]interface{}, len(metrics))
	for i, v := range metrics {
		reqBody[i] = v
	}
	url := getURLForListAndPoints(c)
	query, _ := gophercloud.BuildQueryString(opts)
	url +=query.String()
	_, res.Err = c.Post(url, reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	b := res.Body.(interface{})
	var metricList MetricListData
	err := mapstructure.Decode(b, &metricList)
	return metricList, err
}

// GetDataForListByResolution retrieve data against a list of metrics and specified resolution, for the specified tenant associated with RackspaceMetrics.
func GetDataForListByResolution(c *gophercloud.ServiceClient, opts QueryParams, metrics ...string) (MetricListData, error) {
	var res GetResult

	reqBody := make([]interface{}, len(metrics))
	for i, v := range metrics {
		reqBody[i] = v
	}
	url := getURLForListAndResolution(c)
	query, _ := gophercloud.BuildQueryString(opts)
	url +=query.String()
	_, res.Err = c.Post(url, reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	b := res.Body.(interface{})
	var metricList MetricListData
	err := mapstructure.Decode(b, &metricList)
	return metricList, err
}

// GetDataByPoints retrieve metric data by points, for the specified tenant associated with RackspaceMetrics.
func GetDataByPoints(c *gophercloud.ServiceClient, metric string, opts QueryParams) GetResult {
	var res GetResult

	url := getURLForPoints(c, metric)
	query, _ := gophercloud.BuildQueryString(opts)
	url +=query.String()
	_, res.Err = c.Get(url, &res.Body, nil)
	return res
}

// GetDataByPoints retrieve metric data by resolution, for the specified tenant associated with RackspaceMetrics.
func GetDataByResolution(c *gophercloud.ServiceClient, metric string, opts QueryParams) GetResult {
	var res GetResult

	url := getURLForResolution(c, metric)
	query, _ := gophercloud.BuildQueryString(opts)
	url +=query.String()
	_, res.Err = c.Get(url, &res.Body, nil)
	return res
}

// SearchMetric retrieves a list of available metrics for the specified tenant associated with RackspaceMetrics.
func SearchMetric(c *gophercloud.ServiceClient, opts QueryParams) ([]Metric, error) {
	var res GetResult

	url := getSearchURL(c)
	query, _ := gophercloud.BuildQueryString(opts)
	url +=query.String()

	_, res.Err = c.Get(url, &res.Body, nil)
	b := res.Body.(interface{})
	var metrics []Metric
	err := mapstructure.Decode(b, &metrics)
	return metrics, err
}

//GetEvents retrieves a list of events for the specified tenant associated with RackspaceMetrics.
func GetEvents(c *gophercloud.ServiceClient, opts QueryParams) ([]Event, error) {
	var res GetResult

	url := getEventURL(c)
	query, _ := gophercloud.BuildQueryString(opts)
	url +=query.String()

	_, res.Err = c.Get(url, &res.Body, nil)
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