package query

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
)

// GetResult represents the result of a Get operation.
type GetResult struct {
	gophercloud.Result
}

// MetricListData represents the result of query for multiple metrics.
type MetricListData struct {
	Metrics []MetricList `mapstructure:"metrics"`
}

// MetricList is a part of MetricListData.
type MetricList struct {
	Unit   string `mapstructure:"unit"`
	Metric string `mapstructure:"metric"`
	Data   []Value `mapstructure:"data"`
	Type   string `mapstructure:"type"`
}

// []Metric is the result for metric search query.
type Metric struct {
	Metric string `mapstructure:"metric"`
}

// MetricListData represents the result of query for a metric.
type MetricData struct {
	Unit     string `mapstructure:"unit"`
	Values   []Value `mapstructure:"values"`
	MetaData MetaData `mapstructure:"metadata"`
}

// Value is the metric value. It is also a part of MetricListData and MetricData.
type Value struct {
	NumPoints int32 `mapstructure:"numPoints"`
	TimeStamp int64 `mapstructure:"timestamp"`
	Average   float64 `mapstructure:"average"`
	Sum       float64 `mapstructure:"sum"`
	Max       float64 `mapstructure:"max"`
	Min       float64 `mapstructure:"min"`
	Latest    float64 `mapstructure:"latest"`
	Variance  float64 `mapstructure:"variance"`
}

//Metadata is the additional info about MetricData hence, it is a part of MetricData.
type MetaData struct {
	Limit     string `mapstructure:"limit"`
	Next_Href string `mapstructure:"next_href"`
	Count     int32 `mapstructure:"count"`
	Marker    string `mapstructure:"marker"`
}

//Event is mostly used to get specific events so as to annotate graphs (that's one use-case).
type Event struct {
	What string `mapstructure:"what"`
	When int64 `mapstructure:"when"`
	Data string `mapstructure:"data"`
	Tags string `mapstructure:"tags"`
}

//Limits is used to tell rate-limits for metrics query end-point.
type Limits struct {
	Rate []Rate `mapstructure:"rate"`
}

//Rate is a part of Limits
type Rate struct {
    Limit []Limit `mapstructure:"limit"`
	Regex string `mapstructure:"regex"`
	URI string `mapstructure:"uri"`
}

//Limit is a part of Rate
type Limit struct {
	Next_Available string `mapstructure:"next-available"`
	Remaining int32 `mapstructure:"remaining"`
	Unit string `mapstructure:"unit"`
	Value int32 `mapstructure:"value"`
	Verb string `mapstructure:"verb"`
}

// ExtractMetadata is a function that takes a GetResult (of type *http.Response)
// and returns the custom metatdata associated with the account.
func (gr GetResult) ExtractMetadata() (MetaData, error) {
	var res MetaData
	b := gr.Body.(map[string]interface{})
	err := mapstructure.Decode(b["metadata"], &res)
	return res, err
}

func (gr GetResult) GetValues() ([]Value, error) {
	var res []Value
	b := gr.Body.(map[string]interface{})
	err := mapstructure.Decode(b["values"], &res)
	return res, err
}

func (gr GetResult) GetMetricsData() (MetricData, error) {
	var res MetricData
	b := gr.Body.(interface{})
	err := mapstructure.Decode(b, &res)
	return res, err
}
