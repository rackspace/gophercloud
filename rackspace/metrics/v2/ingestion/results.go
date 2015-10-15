package ingestion
import "github.com/rackspace/gophercloud"

// PostResult represents the result of a Post operation.
type PostResult struct {
	gophercloud.Result
}

//Metric data for ingest payload.
type MetricData struct {
	CollectionTime int64 `mapstructure:"collectionTime"`
	TtlInSeconds   int64 `mapstructure:"ttlInSeconds"`
	MetricValue    float64 `mapstructure:"metricValue"`
	MetricName     string  `mapstructure:"metricName"`
}

//Aggregated metric data for ingest payload.
type AggregatedMetricData struct {
	TenantId  string `mapstructure:"tenantId"`
	Timestamp int64 `mapstructure:"timestamp"`
	Counters  []Counter `mapstructure:"counters"`
	Timers    []Timer `mapstructure:"timers"`
	Sets      []Set `mapstructure:"sets"`
	Gauges    []Gauge `mapstructure:"gauges"`
}

//Counter structure for aggregated metrics payload.
type Counter struct {
	Name  string `mapstructure:"name"`
	Rate  float64 `mapstructure:"rate"`
	Value float64 `mapstructure:"value"`
}

//Timer structure for aggregated metrics payload.
type Timer struct {
	Name        string `mapstructure:"name"`
	Count       int64 `mapstructure:"count"`
	Rate        float64 `mapstructure:"rate"`
	Min         int64 `mapstructure:"min"`
	Max         int64 `mapstructure:"max"`
	Sum         float64 `mapstructure:"sum"`
	Average     float64 `mapstructure:"avg"`
	Std         float64 `mapstructure:"std"`
	Median      int64 `mapstructure:"median"`
	Histograms   []Histogram `mapstructure:"histogram"`
	Percentiles []Percentile `mapstructure:"percentiles"`
}

//Histogram structure for timers payload.
type Histogram struct {
	Bin       string
	Frequency int64
}

//Percentile structure for timers payload.
type Percentile struct {
	Key   int32
	Value Value
}
// Part of percentile
type Value struct {
	Max     int64 `mapstructure:"max"`
	Sum     float64 `mapstructure:"sum"`
	Average float64 `mapstructure:"avg"`
}

//Gauge structure for aggregated metrics payload.
type Gauge struct {
	Name  string `mapstructure:"name"`
	Value float64 `mapstructure:"value"`
}

//Set structure for aggregated metrics payload.
type Set struct {
	Name   string `mapstructure:"name"`
	Values []string `mapstructure:"values"`
}

//Event is mostly used to send specific events so as to annotate graphs (that's one use-case).
type Event struct {
	What string `mapstructure:"what"`
	When int64 `mapstructure:"when"`
	Data string `mapstructure:"data"`
	Tags string `mapstructure:"tags"`
}