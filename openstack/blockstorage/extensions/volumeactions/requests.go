package volumeactions

import (
	"github.com/rackspace/gophercloud"
)

type ConnectorOptsBuilder interface {
	ToConnectorMap() (map[string]interface{}, error)
}

type ConnectorOpts struct {
	IP        string
	Host      string
	Initiator string
	Wwpns     string
	Wwnns     string
	MultiPath bool
	Platform  string
	OSType    string
}

type AttachOpts struct {
	MountPoint   string
	InstanceUUID string
	HostName     string
	Mode         string
}

func Reserve(client *gophercloud.ServiceClient, id string) actionsResult {
	var result actionsResult
	empty_val := make(map[string]interface{})
	reqBody := map[string]interface{}{"os-reserve": empty_val}
	_, result.Err = client.Post(volumeActionsURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return result
}

func (opts ConnectorOpts) ToConnectorMap() (map[string]interface{}, error) {
	c := make(map[string]interface{})

	if opts.IP != "" {
		c["ip"] = opts.IP
	}
	if opts.Host != "" {
		c["host"] = opts.Host
	}
	if opts.Initiator != "" {
		c["initiator"] = opts.Initiator
	}
	if opts.Wwpns != "" {
		c["wwpns"] = opts.Wwpns
	}
	if opts.Wwnns != "" {
		c["wwnns"] = opts.Wwnns
	}
	if opts.Platform != "" {
		c["platform"] = opts.Platform
	}
	if opts.OSType != "" {
		c["os_type"] = opts.OSType
	}
	c["multipath"] = opts.MultiPath
	return map[string]interface{}{"connector": c}, nil
}

func InitializeConnection(client *gophercloud.ServiceClient, id string, opts *ConnectorOpts) actionsResult {
	var result actionsResult
	connector, err := opts.ToConnectorMap()
	if err != nil {
		result.Err = err
		return result
	}
	reqBody := map[string]interface{}{"os-initialize_connection": connector}
	_, result.Err = client.Post(volumeActionsURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		result.Err = err
		return result
	}
	return result
}

func (opts AttachOpts) ToAttachMap() (map[string]interface{}, error) {
	a := make(map[string]interface{})

	if opts.MountPoint != "" {
		a["mountpoint"] = opts.MountPoint
	}
	if opts.Mode != "" {
		a["mode"] = opts.Mode
	}
	if opts.InstanceUUID != "" {
		a["instance_uuid"] = opts.InstanceUUID
	}
	if opts.HostName != "" {
		a["host_name"] = opts.HostName
	}
	return a, nil
}

func Attach(client *gophercloud.ServiceClient, id string, opts *AttachOpts) actionsResult {
	var result actionsResult
	attach, err := opts.ToAttachMap()
	if err != nil {
		result.Err = err
		return result
	}
	reqBody := map[string]interface{}{"os-attach": attach}
	_, result.Err = client.Post(volumeActionsURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		result.Err = err
		return result
	}
	return result

}

func UnReserve(client *gophercloud.ServiceClient, id string) actionsResult {
	var result actionsResult
	empty_val := make(map[string]interface{})
	reqBody := map[string]interface{}{"os-unreserve": empty_val}
	_, result.Err = client.Post(volumeActionsURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return result
}

func TerminateConnection(client *gophercloud.ServiceClient, id string, opts *ConnectorOpts) actionsResult {
	var result actionsResult
	connector, err := opts.ToConnectorMap()
	if err != nil {
		result.Err = err
		return result
	}
	reqBody := map[string]interface{}{"os-terminate_connection": connector}
	_, result.Err = client.Post(volumeActionsURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return result
}

func Detach(client *gophercloud.ServiceClient, id string) actionsResult {
	var result actionsResult
	empty_val := make(map[string]interface{})
	reqBody := map[string]interface{}{"os-detach": empty_val}
	_, result.Err = client.Post(volumeActionsURL(client, id), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return result
}
