package launch

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

// Configuration represents an auto scale group's launch configuration.
type Configuration struct {
	Type string `mapstructure:"type" json:"type"`

	Args Args `mapstructure:"args" json:"args"`
}

// Args represents a launch configuration's arguments.
type Args struct {
	Server map[string]interface{} `mapstructure:"server" json:"server"`

	LoadBalancers []map[string]interface{} `mapstructure:"loadBalancers" json:"loadBalancers"`

	DrainingTimeout int `mapstructure:"draining_timeout" json:"draining_timeout"`
}

type launchResult struct {
	gophercloud.Result
}

// Extract attempts to interpret any launchResult as a Configuration.
func (r launchResult) Extract() (*Configuration, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Configuration Configuration `mapstructure:"launchConfiguration"`
	}

	err := mapstructure.Decode(r.Body, &response)

	if err != nil {
		return nil, err
	}

	return &response.Configuration, nil
}

// GetResult temporarily contains the response from a Get call.
type GetResult struct {
	launchResult
}
