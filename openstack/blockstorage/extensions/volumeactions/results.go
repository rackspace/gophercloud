package volumeactions

import (
	"github.com/rackspace/gophercloud"

	"github.com/mitchellh/mapstructure"
)

// We need a local result type, just declare one here and use the generic
// gophercloud.Result
type actionsResult struct {
	gophercloud.Result
}

func (r actionsResult) Extract() (map[string]interface{}, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	var res map[string]interface{}
	err := mapstructure.Decode(r.Body, &res)
	return res, err
}
