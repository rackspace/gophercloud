package ingestion
import "fmt"

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
