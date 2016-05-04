package launch

import "github.com/rackspace/gophercloud"

// Get requests the details of a given auto scale group's launch configuration.
func Get(client *gophercloud.ServiceClient, groupID string) GetResult {
	var result GetResult

	_, result.Err = client.Get(getURL(client, groupID), &result.Body, nil)

	return result
}

// UpdateOptsBuilder is the interface responsible for generating the map
// structure for producing JSON for an Update operation.
type UpdateOptsBuilder interface {
	ToLaunchUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents the options for updating a launch configuration.
//
// Update operations completely replace the configuration being updated. Empty
// values in the update are accepted and overwrite previously specified
// parameters.
type UpdateOpts struct {
	Configuration
}

// ToLaunchUpdateMap converts an UpdateOpts struct into a map for use as the
// request body in an Update request.
func (opts UpdateOpts) ToLaunchUpdateMap() (map[string]interface{}, error) {
	if err := opts.Configuration.Validate(); err != nil {
		return nil, err
	}

	config := make(map[string]interface{})

	config["type"] = opts.Type
	config["args"] = opts.Args

	return config, nil
}

// Update requests the given configuration be updated.
func Update(client *gophercloud.ServiceClient, groupID string, opts UpdateOptsBuilder) UpdateResult {
	var result UpdateResult

	url := updateURL(client, groupID)
	reqBody, err := opts.ToLaunchUpdateMap()

	if err != nil {
		result.Err = err
		return result
	}

	_, result.Err = client.Put(url, reqBody, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return result
}
