package query

import (
	"testing"
	"net/http"
	"fmt"
	fake "github.com/rackspace/gophercloud/testhelper/client"
	th "github.com/rackspace/gophercloud/testhelper"
)

const DUMMY_METRIC = "rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count"

const DUMMY_METRIC_FOR_LIST_OF_METRICS_BY_POINTS = `{
    "metrics": [
        {
            "unit": "unknown",
            "metric": "rollup00.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
            "data": [
                {
                    "numPoints": 2454,
                    "timestamp": 1444435200000,
                    "average": 26603334416.81
                },
                {
                    "numPoints": 2361,
                    "timestamp": 1444521600000,
                    "average": 28768126895.46
                }
            ],
            "type": "number"
        },
        {
            "unit": "unknown",
            "metric": "rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
            "data": [
                {
                    "numPoints": 2304,
                    "timestamp": 1444435200000,
                    "average": 23165844830.65
                },
                {
                    "numPoints": 2116,
                    "timestamp": 1444521600000,
                    "average": 25393982846.12
                }
            ],
            "type": "number"
        }
    ]
}`

//Test results for a list of metrics input queried by points.
func TestGetDataForListByPoints(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/views", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, DUMMY_METRIC_FOR_LIST_OF_METRICS_BY_POINTS)
	})

	// Expected metric data when queried using list of metrics.
	expectedMetrics := MetricListData{
		Metrics: []MetricList{
			MetricList{
				Unit: "unknown",
				Metric: "rollup00.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
				Data: []Value{
					Value{
						NumPoints: 2454,
						TimeStamp: 1444435200000,
						Average: 26603334416.81,
					},
					Value{
						NumPoints: 2361,
						TimeStamp: 1444521600000,
						Average: 28768126895.46,
					},
				},
				Type: "number",
			},
			MetricList{
				Unit: "unknown",
				Metric: "rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
				Data: []Value{
					Value{
						NumPoints: 2304,
						TimeStamp: 1444435200000,
						Average: 23165844830.65,
					},
					Value{
						NumPoints: 2116,
						TimeStamp: 1444521600000,
						Average: 25393982846.12,
					},
				},
				Type: "number",
			},
		},
	}

	queryParams := QueryParams{
		From: 1444354736000,
		To: 1444602340000,
		Points: 3,
	}
	// Actual metric data when queried using list of metrics.
	actualMetrics, err := GetDataForListByPoints(fake.ServiceClient(), queryParams,
		"rollup00.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
		"rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetrics, actualMetrics)
}

const DUMMY_METRIC_FOR_LIST_OF_METRICS_BY_RESOLUTION = `{
    "metrics": [
        {
            "unit": "unknown",
            "metric": "rollup02.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
            "data": [
                {
                    "numPoints": 2431,
                    "timestamp": 1444348800000,
                    "average": 24496343044.07
                },
                {
                    "numPoints": 2454,
                    "timestamp": 1444435200000,
                    "average": 26603334416.81
                },
                {
                    "numPoints": 2361,
                    "timestamp": 1444521600000,
                    "average": 28768126895.46
                }
            ],
            "type": "number"
        },
        {
            "unit": "unknown",
            "metric": "rollup03.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
            "data": [
                {
                    "numPoints": 2241,
                    "timestamp": 1444348800000,
                    "average": 20980428271.97
                },
                {
                    "numPoints": 2304,
                    "timestamp": 1444435200000,
                    "average": 23165844830.65
                },
                {
                    "numPoints": 2116,
                    "timestamp": 1444521600000,
                    "average": 25393982846.12
                }
            ],
            "type": "number"
        }
    ]
}`

//Test results for a list of metrics input queried by resolution.
func TestGetDataForListByResolution(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/views", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, DUMMY_METRIC_FOR_LIST_OF_METRICS_BY_RESOLUTION)
	})

	// Expected metric data when queried using list of metrics.
	expectedMetrics := MetricListData{
		Metrics: []MetricList{
			MetricList{
				Unit: "unknown",
				Metric: "rollup02.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
				Data: []Value{
					Value{
						NumPoints: 2431,
						TimeStamp: 1444348800000,
						Average: 24496343044.07,
					},
					Value{
						NumPoints: 2454,
						TimeStamp: 1444435200000,
						Average: 26603334416.81,
					},
					Value{
						NumPoints: 2361,
						TimeStamp: 1444521600000,
						Average: 28768126895.46,
					},
				},
				Type: "number",
			},
			MetricList{
				Unit: "unknown",
				Metric: "rollup03.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
				Data: []Value{
					Value{
						NumPoints: 2241,
						TimeStamp: 1444348800000,
						Average: 20980428271.97,
					},
					Value{
						NumPoints: 2304,
						TimeStamp: 1444435200000,
						Average: 23165844830.65,
					},
					Value{
						NumPoints: 2116,
						TimeStamp: 1444521600000,
						Average: 25393982846.12,
					},
				},
				Type: "number",
			},
		},
	}

	queryParams := QueryParams{
		From: 1444354736000,
		To: 1444602340000,
		Resolution: MIN1440,
	}

	// Actual metric data when queried using list of metrics.
	actualMetrics, err := GetDataForListByResolution(fake.ServiceClient(), queryParams,
		"rollup02.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count",
		"rollup03.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetrics, actualMetrics)
}

const DUMMY_METRIC_DATA_FOR_FIVE_POINTS = `{
    "metadata": {
        "count": 5,
        "limit": null,
        "marker": null,
        "next_href": null
    },
    "unit": "unknown",
    "values": [
        {
            "average": 20015507706.5,
            "numPoints": 32,
            "timestamp": 1444353600000
        },
        {
            "average": 20045095111.91,
            "numPoints": 32,
            "timestamp": 1444354800000
        },
        {
            "average": 20077360883.33,
            "numPoints": 33,
            "timestamp": 1444356000000
        },
        {
            "average": 20110963258.48,
            "numPoints": 25,
            "timestamp": 1444357200000
        },
        {
            "average": 25588472512.28,
            "numPoints": 330,
            "timestamp": 1444357300000
        }
    ]
}`

const DUMMY_METRIC_DATA_FOR_TWO_POINTS = `{
    "metadata": {
        "count": 2,
        "limit": null,
        "marker": null,
        "next_href": null
    },
    "unit": "unknown",
    "values": [
        {
            "average": 20015507706.5,
            "numPoints": 32,
            "timestamp": 1444353600000
        },
        {
            "average": 20045095111.91,
            "numPoints": 32,
            "timestamp": 1444354800000
        }
    ]
}`

const DUMMY_METRIC_DATA_FOR_TWO_POINTS_WITH_MAX_SELECT_PARAM = `{
    "metadata": {
        "count": 2,
        "limit": null,
        "marker": null,
        "next_href": null
    },
    "unit": "unknown",
    "values": [
        {
            "max": 20015507706.5,
            "timestamp": 1444353600000
        },
        {
            "max": 20045095111.91,
            "timestamp": 1444354800000
        }
    ]
}`

const DUMMY_METRIC_DATA_FOR_TWO_POINTS_WITH_MIN_SELECT_PARAM = `{
    "metadata": {
        "count": 2,
        "limit": null,
        "marker": null,
        "next_href": null
    },
    "unit": "unknown",
    "values": [
        {
            "min": 20015507706.5,
            "timestamp": 1444353600000
        },
        {
            "min": 20045095111.91,
            "timestamp": 1444354800000
        }
    ]
}`

//Test results for a metric input queried by points.
func TestGetDataByPoints(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/views/"+DUMMY_METRIC, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.URL.Query().Get("select") == "" {
			if r.URL.Query().Get("points") == "5" {
				fmt.Fprintf(w, DUMMY_METRIC_DATA_FOR_FIVE_POINTS)
			} else if r.URL.Query().Get("points") == "2" {
				fmt.Fprintf(w, DUMMY_METRIC_DATA_FOR_TWO_POINTS)
			}
		} else {
			if r.URL.Query().Get("select") == "max" {
				fmt.Fprintf(w, DUMMY_METRIC_DATA_FOR_TWO_POINTS_WITH_MAX_SELECT_PARAM)
			} else if r.URL.Query().Get("select") == "min" {
				fmt.Fprintf(w, DUMMY_METRIC_DATA_FOR_TWO_POINTS_WITH_MIN_SELECT_PARAM)
			}
		}
	})

	//Expected metric data for 5 points.
	expectedMetricDataFor5Points := MetricData{
		MetaData: MetaData{
			Count: 5,
		},
		Unit: "unknown",
		Values: []Value{
			Value{
				Average: 20015507706.50,
				NumPoints: 32,
				TimeStamp: 1444353600000,
			},
			Value{
				Average: 20045095111.91,
				NumPoints: 32,
				TimeStamp: 1444354800000,
			},
			Value{
				Average: 20077360883.33,
				NumPoints: 33,
				TimeStamp: 1444356000000,
			},
			Value{
				Average: 20110963258.48,
				NumPoints: 25,
				TimeStamp: 1444357200000,
			},
			Value{
				Average: 25588472512.28,
				NumPoints: 330,
				TimeStamp: 1444357300000,
			},
		},
	}

	//Expected metric values (part of metric data) for 5 points.
	expectedValuesFor5Points := []Value{
		Value{
			Average: 20015507706.50,
			NumPoints: 32,
			TimeStamp: 1444353600000,
		},
		Value{
			Average: 20045095111.91,
			NumPoints: 32,
			TimeStamp: 1444354800000,
		},
		Value{
			Average: 20077360883.33,
			NumPoints: 33,
			TimeStamp: 1444356000000,
		},
		Value{
			Average: 20110963258.48,
			NumPoints: 25,
			TimeStamp: 1444357200000,
		},
		Value{
			Average: 25588472512.28,
			NumPoints: 330,
			TimeStamp: 1444357300000,
		},
	}

	//Expected metric metadata (part of metric data) for 5 points.
	expectedMetaDataFor5Points := MetaData{
		Count: 5,
	}

	//Expected metric data for 2 points.
	expectedMetricDataFor2Points := MetricData{
		MetaData: MetaData{
			Count: 2,
		},
		Unit: "unknown",
		Values: []Value{
			Value{
				Average: 20015507706.50,
				NumPoints: 32,
				TimeStamp: 1444353600000,
			},
			Value{
				Average: 20045095111.91,
				NumPoints: 32,
				TimeStamp: 1444354800000,
			},
		},
	}

	//Expected Metric data for 2 points with "max" as select param.
	expectedMetricDataFor2PointsWithMaxSelectParam := MetricData{
		MetaData: MetaData{
			Count: 2,
		},
		Unit: "unknown",
		Values: []Value{
			Value{
				Max: 20015507706.50,
				TimeStamp: 1444353600000,
			},
			Value{
				Max: 20045095111.91,
				TimeStamp: 1444354800000,
			},
		},
	}

	//Expected Metric data for 2 points with "max" as select param.
	expectedMetricDataFor2PointsWithMinSelectParam := MetricData{
		MetaData: MetaData{
			Count: 2,
		},
		Unit: "unknown",
		Values: []Value{
			Value{
				Min: 20015507706.50,
				TimeStamp: 1444353600000,
			},
			Value{
				Min: 20045095111.91,
				TimeStamp: 1444354800000,
			},
		},
	}

	queryParamsFor5Points := QueryParams{
		From: 1444354736000,
		To: 1444602340000,
		Points: 5,
	}

	//Actual Metric data for 5 points.
	actualMetricDataFor5Points, err := GetDataByPoints(fake.ServiceClient(), DUMMY_METRIC, queryParamsFor5Points).GetMetricsData()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetricDataFor5Points, actualMetricDataFor5Points)

	//Actual metric values (part of metric data) for 5 points.
	actualValuesFor5Points, err := GetDataByPoints(fake.ServiceClient(), DUMMY_METRIC, queryParamsFor5Points).GetValues()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedValuesFor5Points, actualValuesFor5Points)

	//Actual metadata (part of metric data) for 5 points.
	actualMetaDataFor5Points, err := GetDataByPoints(fake.ServiceClient(), DUMMY_METRIC, queryParamsFor5Points).ExtractMetadata()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetaDataFor5Points, actualMetaDataFor5Points)

	queryParamsFor2Points := QueryParams{
		From: 1444354736000,
		To: 1444602340000,
		Points: 2,
	}

	//Actual Metric data for 2 points.
	actualMetricDataFor2Points, err := GetDataByPoints(fake.ServiceClient(), DUMMY_METRIC, queryParamsFor2Points).GetMetricsData()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetricDataFor2Points, actualMetricDataFor2Points)

	queryParamsFor2PointsWithMax := QueryParams{
		From: 1444354736000,
		To: 1444602340000,
		Points:2,
		Select: []string{MAX},
	}

	//Actual Metric data for 2 points with "max" as select param.
	actualMetricDataFor2PointsWithMaxSelectParam, err := GetDataByPoints(fake.ServiceClient(), DUMMY_METRIC, queryParamsFor2PointsWithMax).GetMetricsData()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetricDataFor2PointsWithMaxSelectParam, actualMetricDataFor2PointsWithMaxSelectParam)

	queryParamsFor2PointsWithMin := QueryParams{
		From: 1444354736000,
		To: 1444602340000,
		Points:2,
		Select: []string{MIN},
	}

	//Actual Metric data for 2 points with "min" as select param.
	actualMetricDataFor2PointsWithMinSelectParam, err := GetDataByPoints(fake.ServiceClient(), DUMMY_METRIC, queryParamsFor2PointsWithMin).GetMetricsData()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetricDataFor2PointsWithMinSelectParam, actualMetricDataFor2PointsWithMinSelectParam)
}

const DUMMY_METRIC_DATA_FOR_MIN1440 = `{
    "metadata": {
        "count": 3,
        "limit": null,
        "marker": null,
        "next_href": null
    },
    "unit": "unknown",
    "values": [
        {
            "average": 209804282719.7,
            "numPoints": 2241,
            "timestamp": 1444348800000
        },
        {
            "average": 231658448306.51,
            "numPoints": 2304,
            "timestamp": 1444435200000
        },
        {
            "average": 253939828461.21,
            "numPoints": 2116,
            "timestamp": 1444521600000
        }
    ]
}`

const DUMMY_METRIC_DATA_FOR_MIN1440_WITH_MIN_SELECT_PARAM = `{
    "metadata": {
        "count": 3,
        "limit": null,
        "marker": null,
        "next_href": null
    },
    "unit": "unknown",
    "values": [
        {
            "min": 209804282719.7,
            "timestamp": 1444348800000
        },
        {
            "min": 231658448306.51,
            "timestamp": 1444435200000
        },
        {
            "min": 253939828461.21,
            "timestamp": 1444521600000
        }
    ]
}`

//Test results for a metric input queried by resolution.
func TestGetDataByResolution(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/views/"+DUMMY_METRIC, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.URL.Query().Get("select") == "" {
			if r.URL.Query().Get("resolution") == "min1440" {
				fmt.Fprintf(w, DUMMY_METRIC_DATA_FOR_MIN1440)
			}
		} else {
			fmt.Fprintf(w, DUMMY_METRIC_DATA_FOR_MIN1440_WITH_MIN_SELECT_PARAM)
		}
	})

	//Expected metric data for Min1440.
	expectedMetricDataForMin1440 := MetricData{
		MetaData: MetaData{
			Count: 3,
		},
		Unit: "unknown",
		Values: []Value{
			Value{
				Average: 209804282719.70,
				NumPoints: 2241,
				TimeStamp: 1444348800000,
			},
			Value{
				Average: 231658448306.51,
				NumPoints: 2304,
				TimeStamp: 1444435200000,
			},
			Value{
				Average: 253939828461.21,
				NumPoints: 2116,
				TimeStamp: 1444521600000,
			},
		},
	}

	//Expected metric data for Min1440.
	expectedMetricDataForMin1440WithMinSelectParam := MetricData{
		MetaData: MetaData{
			Count: 3,
		},
		Unit: "unknown",
		Values: []Value{
			Value{
				Min: 209804282719.70,
				TimeStamp: 1444348800000,
			},
			Value{
				Min: 231658448306.51,
				TimeStamp: 1444435200000,
			},
			Value{
				Min: 253939828461.21,
				TimeStamp: 1444521600000,
			},
		},
	}

	queryParamsForMin1440Resolution := QueryParams{
		From: 1444354736000,
		To: 1444355736000,
		Resolution: MIN1440,
	}

	//Actual Metric data for Min1440.
	actualMetricDataFor5Points, err := GetDataByResolution(fake.ServiceClient(), DUMMY_METRIC, queryParamsForMin1440Resolution).GetMetricsData()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetricDataForMin1440, actualMetricDataFor5Points)

	queryParamsForMin1440ResolutionWithMinSelect := QueryParams{
		From: 1444354736000,
		To: 1444355736000,
		Resolution: MIN1440,
		Select:[]string{MIN},
	}

	//Actual Metric data for Min1440 with "max" as select param..
	actualMetricDataFor5PointsWithMinSelectParam, err := GetDataByResolution(fake.ServiceClient(), DUMMY_METRIC, queryParamsForMin1440ResolutionWithMinSelect).GetMetricsData()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetricDataForMin1440WithMinSelectParam, actualMetricDataFor5PointsWithMinSelectParam)
}

//Test results metric search using wildcards.
func TestSearchMetric(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/metrics/search", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.URL.Query().Get("query") == "*.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count" {
			fmt.Fprintf(w, `[
    {
        "metric": "rollup00.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count"
    },
    {
        "metric": "rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count"
    },
    {
        "metric": "rollup02.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count"
    },
    {
        "metric": "rollup03.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count"
    }
]`)
		}
	})

	// Expected events data when no tag is specified.
	expectedMetrics := []Metric{
		Metric{
			Metric:   "rollup00.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count",
		},
		Metric{
			Metric:   "rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count",
		},
		Metric{
			Metric:   "rollup02.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count",
		},
		Metric{
			Metric:   "rollup03.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count",
		},
	}

	queryParamsForSearch := QueryParams{
		Query: "*.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.ShardStatePuller.Stats.count",
	}

	// Actual metrics received for search query.
	actualMetrics, err := SearchMetric(fake.ServiceClient(), queryParamsForSearch)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedMetrics, actualMetrics)
}

const DUMMY_LIMITS = `{
  "limits": {
    "rate": [
      {
        "limit": [
          {
            "next-available": "2015-10-13T16:42:18.899Z",
            "remaining": 993,
            "unit": "MINUTE",
            "value": 1000,
            "verb": "ALL"
          }
        ],
        "regex": "/v[0-9.]+/((hybrid:)?[0-9]+)/.+",
        "uri": "/version/tenantId/*"
      }
    ]
  }
}`

//Test results for events.
func TestGetEvents(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/events/getEvents", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.URL.Query().Get("tags") == "" {
			fmt.Fprintf(w, `[
    {
        "what": "Test Event 1",
        "when": 1444354736003,
        "data": "Test Data 1",
        "tags": "Restart"
    },
    {
        "what": "Test Event 2",
        "when": 1444354736104,
        "data": "Test Data 2",
        "tags": "Shutdown"
    }
]`)
		} else if r.URL.Query().Get("tags") == "Restart" {
			fmt.Fprintf(w, `[
    {
        "what": "Test Event 1",
        "when": 1444354736003,
        "data": "Test Data 1",
        "tags": "Restart"
    }
]`)
		}
	})

	// Expected events data when no tag is specified.
	expectedEvents := []Event{
		Event{
			What:   "Test Event 1",
			When: 1444354736003,
			Data: "Test Data 1",
			Tags: "Restart",
		},
		Event{
			What:   "Test Event 2",
			When: 1444354736104,
			Data: "Test Data 2",
			Tags: "Shutdown",
		},
	}

	// Expected events data when tag is specified.
	expectedEventsWithTags := []Event{
		Event{
			What:   "Test Event 1",
			When: 1444354736003,
			Data: "Test Data 1",
			Tags: "Restart",
		},
	}

	queryParamsForEvents := QueryParams{
		From: 1444354736000,
		Until: 1444355736000,
	}

	// Actual events data received when no tag is specified.
	actualEvents, err := GetEvents(fake.ServiceClient(), queryParamsForEvents)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedEvents, actualEvents)

	queryParamsForEventsWithTags := QueryParams{
		From: 1444354736000,
		Until: 1444355736000,
		Tags: "Restart",
	}

	// Actual events data received when tag is specified.
	actualEventsWithTags, err := GetEvents(fake.ServiceClient(), queryParamsForEventsWithTags)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedEventsWithTags, actualEventsWithTags)
}

//Test results for rate-limits of query end-point.
func TestGetLimits(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/limits", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, DUMMY_LIMITS)
	})

	// Expected Limits.
	expectedLimits := Limits{
		Rate: []Rate{
			Rate{
				Limit: []Limit{
					Limit{
						Next_Available: "2015-10-13T16:42:18.899Z",
						Remaining: 993,
						Unit: "MINUTE",
						Value: 1000,
						Verb: "ALL",
					},
				},
				Regex: "/v[0-9.]+/((hybrid:)?[0-9]+)/.+",
				URI: "/version/tenantId/*",
			},
		},
	}

	// Actual Limits.
	actualLimits, err := GetLimits(fake.ServiceClient())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedLimits, actualLimits)
}