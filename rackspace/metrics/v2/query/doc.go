// Package query provides access to query end-points associated with a Rackspace Metrics configuration.
// Using this package, user can query for METRICS data either by number of points
// or by resolution. Additional parameters like "SUM", "MIN", "MAX", "LATEST", "VARIANCE" etc, can be send
// in order to get pre-aggregated results.
// Query package also gives you access to query for specific EVENTS, which are stored against the
// associated tenant. You can also discover the metrics associated with your tenant using wild cards.
package query
