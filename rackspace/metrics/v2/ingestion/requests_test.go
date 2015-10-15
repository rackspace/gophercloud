package ingestion

import (
	"testing"
	"net/http"
	"fmt"
	fake "github.com/rackspace/gophercloud/testhelper/client"
	th "github.com/rackspace/gophercloud/testhelper"
)

//Test for sending metrics.
func TestSendMetrics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `[
    {
        "collectionTime": 1376509892612,
        "ttlInSeconds": 172800,
        "metricValue": 66,
        "metricName": "example.metric.one"
    },
    {
        "collectionTime": 1376509892612,
        "ttlInSeconds": 172800,
        "metricValue": 66,
        "metricName": "example.metric.two"
    },
    {
        "collectionTime": 1376509892612,
        "ttlInSeconds": 172800,
        "metricValue": 66,
        "metricName": "example.metric.three"
    }
]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	// Actual metric data send list of metrics.
	actualMetrics := []MetricData{
		MetricData{
			CollectionTime: 1376509892612,
			TtlInSeconds: 172800,
			MetricValue: 66,
			MetricName: "example.metric.one",
		},
		MetricData{
			CollectionTime: 1376509892612,
			TtlInSeconds: 172800,
			MetricValue: 66,
			MetricName: "example.metric.two",
		},
		MetricData{
			CollectionTime: 1376509892612,
			TtlInSeconds: 172800,
			MetricValue: 66,
			MetricName: "example.metric.three",
		},
	}

	SendMetrics(fake.ServiceClient(), actualMetrics)
}

func TestSendAggregatedMetrics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/ingest/aggregated", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `{
    "tenantId": "5405533",
    "timestamp": 1429567619000,
    "counters": [
        {
            "name": "sample_counter",
            "rate": 1,
            "value": 32
        },
        {
            "name": "sample_another_counter",
            "rate": 1,
            "value": 4424
        }
    ],
    "timers": [
			{
				"name": "timer_name",
				"count": 32,
				"rate": 2.3,
				"min": 1,
				"max": 5,
				"sum": 21,
				"avg": 2.1,
				"median": 3,
				"std": 1.01,
				"percentiles": {
					"50": {
						"avg": 121.72972972972973,
						"max": 241,
						"sum": 4504
					},
					"75": {
						"avg": 185.94642857142858,
						"max": 372,
						"sum": 10413
					}
				},
				"histogram": {
					"bin_50": 3,
					"bin_100": 5,
					"bin_inf": 2
				}
			}
		],
		"gauges": [
        {
            "name": "gauge_name",
            "value": 42
        },
        {
            "name": "another_gauge",
            "value": 4343
        }
    ],
    "sets": [
            {
             "name": "set_name",
             "values": ["foo", "bar", "baz"]
            },
           {
            "name": "another_set",
            "values": ["boo", "far", "zab"]
           }
          ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	aggregatedMetrics := AggregatedMetricData{
		TenantId: "5405533",
		Timestamp: 1429567619000,
		Counters : []Counter{
			Counter{
				Name: "sample_counter",
				Rate:1,
				Value:32,
			},
			Counter{
				Name: "sample_another_counter",
				Rate:1,
				Value:4424,
			},
		},
		Timers:  []Timer{
			Timer{
				Name: "timer_name",
				Count:32,
				Rate:2.3,
				Min:1,
				Max:5,
				Sum:21,
				Average:2.1,
				Median:3,
				Std:1.01,
				Percentiles: []Percentile{
					Percentile{
						Key: 50,
						Value: Value{
							Average: 121.72972972972973,
							Max: 241,
							Sum: 4504,
						},
					},
					Percentile{
						Key: 75,
						Value: Value{
							Average: 185.94642857142858,
							Max: 372,
							Sum: 10413,
						},
					},
				},
				Histograms: []Histogram{
					Histogram{
						Bin:"50",
						Frequency:3,
					},
					Histogram{
						Bin:"100",
						Frequency:5,
					},
					Histogram{
						Bin:"inf",
						Frequency:2,
					},
				},
			},
		},
		Gauges : []Gauge{
			Gauge{
				Name: "gauge_name",
				Value:42,
			},
			Gauge{
				Name: "another_gauge",
				Value:4343,
			},
		},
		Sets : []Set{
			Set{
				Name: "set_name",
				Values: []string{"foo", "bar", "baz"},
			},
			Set{
				Name: "another_set",
				Values: []string{"boo", "far", "zab"},
			},
		},
	}

	SendAggregatedMetrics(fake.ServiceClient(), aggregatedMetrics)
}

func TestSendAggregatedCounters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/ingest/aggregated", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `{
    "tenantId": "5405533",
    "timestamp": 1429567619000,
    "counters": [
        {
            "name": "sample_counter",
            "rate": 1,
            "value": 32
        },
        {
            "name": "sample_another_counter",
            "rate": 1,
            "value": 4424
        }
    ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	aggregatedCounters := []Counter{
		Counter{
			Name: "sample_counter",
			Rate:1,
			Value:32,
		},
		Counter{
			Name: "sample_another_counter",
			Rate:1,
			Value:4424,
		},
	}

	SendAggregatedCounters(fake.ServiceClient(), "5405533", 1429567619000, aggregatedCounters)
}

func TestSendAggregatedTimers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/ingest/aggregated", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		//fmt.Println(r.Body)
		th.TestJSONRequest(t, r,
			`
		{
		"tenantId": "5405533",
		"timestamp": 1433451699000,
		"timers": [
			{
				"name": "timer_name",
				"count": 32,
				"rate": 2.3,
				"min": 1,
				"max": 5,
				"sum": 21,
				"avg": 2.1,
				"median": 3,
				"std": 1.01,
				"percentiles": {
					"50": {
						"avg": 121.72972972972973,
						"max": 241,
						"sum": 4504
					},
					"75": {
						"avg": 185.94642857142858,
						"max": 372,
						"sum": 10413
					}
				},
				"histogram": {
					"bin_50": 3,
					"bin_100": 5,
					"bin_inf": 2
				}
			}
		]
		}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	aggregatedTimers := []Timer{
		Timer{
			Name: "timer_name",
			Count:32,
			Rate:2.3,
			Min:1,
			Max:5,
			Sum:21,
			Average:2.1,
			Median:3,
			Std:1.01,
			Percentiles: []Percentile{
				Percentile{
					Key: 50,
					Value: Value{
						Average: 121.72972972972973,
						Max: 241,
						Sum: 4504,
					},
				},
				Percentile{
					Key: 75,
					Value: Value{
						Average: 185.94642857142858,
						Max: 372,
						Sum: 10413,
					},
				},
			},
			Histograms: []Histogram{
				Histogram{
					Bin:"50",
					Frequency:3,
				},
				Histogram{
					Bin:"100",
					Frequency:5,
				},
				Histogram{
					Bin:"inf",
					Frequency:2,
				},
			},
		},
	}

	SendAggregatedTimers(fake.ServiceClient(), "5405533", 1433451699000, aggregatedTimers)
}

func TestSendAggregatedGauges(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/ingest/aggregated", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `{
    "tenantId": "5405533",
    "timestamp": 1433451699000,
    "gauges": [
        {
            "name": "gauge_name",
            "value": 42
        },
        {
            "name": "another_gauge",
            "value": 4343
        }
    ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	aggregatedGauges := []Gauge{
		Gauge{
			Name: "gauge_name",
			Value:42,
		},
		Gauge{
			Name: "another_gauge",
			Value:4343,
		},
	}

	SendAggregatedGauges(fake.ServiceClient(), "5405533", 1433451699000, aggregatedGauges)
}

func TestSendAggregatedSets(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/ingest/aggregated", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `{
  "tenantId": "5405533",
  "timestamp": 1433451699000,
  "sets": [
            {
             "name": "set_name",
             "values": ["foo", "bar", "baz"]
            },
           {
            "name": "another_set",
            "values": ["boo", "far", "zab"]
           }
          ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	aggregatedSets := []Set{
		Set{
			Name: "set_name",
			Values: []string{"foo", "bar", "baz"},
		},
		Set{
			Name: "another_set",
			Values: []string{"boo", "far", "zab"},
		},
	}

	SendAggregatedSets(fake.ServiceClient(), "5405533", 1433451699000, aggregatedSets)
}

func TestSendEvent(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `{
    "what": "Test Event",
    "when": 1441831996000,
    "tags": "Restart",
    "data": "Test Data"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ``)
	})

	event := Event{
		What:"Test Event",
		When:1441831996000,
		Tags:"Restart",
		Data:"Test Data",
	}

	SendEvent(fake.ServiceClient(), event)
}