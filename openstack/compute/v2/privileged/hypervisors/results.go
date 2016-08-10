package hypervisors

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
	"reflect"
)


// Decodes response body. Accepts empty entity pointer and initiates
// this entity with decoded values.
func processResponse(response interface{}, body interface{}) error{
	config := &mapstructure.DecoderConfig{
		DecodeHook: toMapFromString,
		Result:     response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(body)
	if err != nil {
		return err
	}
	return err
}

type hypervisorResult struct {
	gophercloud.Result
}

// GetResult temporarily contains the response from a Get call.
type GetResult struct {
	hypervisorResult
}


// Extract interprets any hypervisorResult as a Hypervisor, if possible.
func (r hypervisorResult) Extract() ([]Hypervisor, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisors []Hypervisor `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, r.Body)
	return response.Hypervisors, err
}

// Extract interprets any hypervisorResult as a HypervisorDetails, if possible.
func (r hypervisorResult) ExtractDetails() ([]HypervisorDetail, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor []HypervisorDetail `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, r.Body)
	return response.Hypervisor, err
}

// Extract interprets any hypervisorResult as a HypervisorDetail, if possible.
func (r hypervisorResult) ExtractDetail() (*HypervisorDetail, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor HypervisorDetail `mapstructure:"hypervisor"`
	}
	err := processResponse(&response, r.Body)
	return &response.Hypervisor, err
}

// Extract interprets any hypervisorResult as a HypervisorServersInfo, if possible.
func (r hypervisorResult) ExtractServersInfo() ([]HypervisorServersInfo, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor []HypervisorServersInfo `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, r.Body)
	return response.Hypervisor, err
}

// Extract interprets any hypervisorResult as a HypervisorUptime, if possible.
func (r hypervisorResult) ExtractUptime() (*HypervisorUptimeInfo, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor HypervisorUptimeInfo `mapstructure:"hypervisor"`
	}
	err := processResponse(&response, r.Body)
	return &response.Hypervisor, err
}

func toMapFromString(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
	if (from == reflect.String) && (to == reflect.Map) {
		return map[string]interface{}{}, nil
	}
	return data, nil
}
