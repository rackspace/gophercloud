package launch

import (
	"encoding/base64"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

// Network represents a reference to a Cloud Network.
type Network struct {
	UUID string `mapstructure:"uuid" json:"uuid"`
}

// Default Rackspace networks.
var (
	PublicNet  = Network{UUID: "00000000-0000-0000-0000-000000000000"}
	ServiceNet = Network{UUID: "11111111-1111-1111-1111-111111111111"}
)

// LoadBalancerType is a string representing the type of load balancer servers
// in a group will be attached to.
type LoadBalancerType string

// Valid load balancer types.
const (
	CloudLoadBalancer LoadBalancerType = "CloudLoadBalancer"
	RackConnectV3     LoadBalancerType = "RackConnectV3"
)

// LoadBalancer represents the details of a load balancer that servers in a
// group will be attached to.
type LoadBalancer struct {
	// Type of load balancer: CloudLoadBalancer or RackConnectV3
	Type LoadBalancerType

	// UUID of the RackConnectV3 load balancer, or an empty string in the case
	// of a Cloud Load Balancer.
	RackConnectUUID string

	// ID of a Cloud Load Balancers load balancer.
	CloudLoadBalancerID int

	// Port on the servers in a group that the load balancer will use.  Will be
	// zero for RackConnectV3 load balancers, where this parameter is not used.
	Port int
}

// Server represents the attributes used to create a new server in a group.
type Server struct {
	// Base name for servers in the group.
	Name string `mapstructure:"name" json:"name"`

	// Flavor of server to be created.
	FlavorRef string `mapstructure:"flavorRef" json:"flavorRef"`

	// ID of the server image used for new servers.
	ImageRef string `mapstructure:"imageRef" json:"imageRef"`

	// Disk Configuration mode: MANUAL, AUTO, or an empty string if no mode has
	// been specified.
	DiskConfig string `mapstructure:"OS-DCF:diskConfig" json:"OS-DCF:diskConfig"`

	// Name of a preexisting keypair injected into new servers, or an empty
	// string if no keypair has been specified.
	KeyName string `mapstructure:"key_name" json:"key_name"`

	// Whether metadata injection through a configuration drive is enabled.
	ConfigDrive bool `mapstructure:"config_drive" json:"config_drive"`

	// User provided configuration data. Base64 decoded.
	UserData []byte `mapstructure:"user_data" json:"user_data"`

	// Networks new servers will be attached to.
	Networks []Network `mapstructure:"networks" json:"networks"`

	// List of file paths and contents injected into new servers.
	Personality servers.Personality `mapstructure:"Personality" json:"Personality"`

	// Additonal metadata associated with new servers.
	Metadata map[string]interface{} `mapstructure:"metadata" json:"metadata"`
}

// ConfigurationType represents a type of launch configuration.
type ConfigurationType string

// Valid launch configuration types.
const (
	LaunchServer ConfigurationType = "launch_server"
)

// Configuration represents an auto scale group's launch configuration.
type Configuration struct {
	// Type for this configuration.
	Type ConfigurationType `mapstructure:"type" json:"type"`

	// Configuration arguments.
	Args Args `mapstructure:"args" json:"args"`
}

// Args represents a launch configuration's arguments.
type Args struct {
	// Attributes used to create new servers in the group.
	Server Server `mapstructure:"server" json:"server"`

	// List of load balancers to which to attach servers.
	LoadBalancers []LoadBalancer `mapstructure:"loadBalancers" json:"loadBalancers"`

	// Number of seconds a deleted node is put into DRAINING mode in attached
	// load balancers before actually being removed.  This will be zero if no
	// timeout as been specified, in which case nodes will not be drained.
	DrainingTimeout int `mapstructure:"draining_timeout" json:"draining_timeout,omitempty"`
}

// mapstructure.DecodeHookFuncType to convert from a map[string]interface{} to a
// servers.File, decoding contents as base64 if possible.
func mapToFile(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	fileType := reflect.TypeOf((*servers.File)(nil))

	if from.Kind() != reflect.Map || to != fileType {
		return data, nil
	}

	fileMap, ok := data.(map[string]interface{})

	if !ok {
		return data, nil
	}

	file := servers.File{}

	// If the path key exists and is a string, set file.Path.
	if p, ok := fileMap["path"].(string); ok {
		file.Path = p
	}

	// If the contents key exists and is a string, decode the string as base64
	// and set file.Contents.  If decoding fails, use the raw bytes.
	if c, ok := fileMap["contents"].(string); ok {
		bytes, err := base64.StdEncoding.DecodeString(c)

		if err == nil {
			file.Contents = bytes
		} else {
			file.Contents = []byte(c)
		}
	}

	return &file, nil
}

// mapstructure.DecodeHookFuncType to convert from a string to []byte decoding
// contents as base64 if possible.
func stringToBytes(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	byteType := reflect.TypeOf(byte(0))
	byteSlice := reflect.SliceOf(byteType)

	if from.Kind() != reflect.String || to != byteSlice {
		return data, nil
	}

	str, ok := data.(string)

	if !ok {
		return data, nil
	}

	bytes, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		return []byte(str), nil
	}

	return bytes, nil
}

// mapstructure.DecodeHookFuncType to convert from a map[string]interface{} to a
// LoadBalancer struct.
func mapToLoadBalancer(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	lbType := reflect.TypeOf((*LoadBalancer)(nil)).Elem()

	if from.Kind() != reflect.Map || to != lbType {
		return data, nil
	}

	lbMap, ok := data.(map[string]interface{})

	if !ok {
		return data, nil
	}

	lb := LoadBalancer{}

	// If the ID is a string, it's a RackConnect V3 UUID, else it should be a
	// Cloud Load Balancers ID, which is an integer.
	if id, ok := lbMap["loadBalancerId"].(string); ok {
		lb.RackConnectUUID = id
	} else if id, ok := lbMap["loadBalancerId"].(float64); ok {
		lb.CloudLoadBalancerID = int(id)
	}

	if p, ok := lbMap["port"].(float64); ok {
		lb.Port = int(p)
	}

	if t, ok := lbMap["type"].(string); ok {
		lb.Type = LoadBalancerType(t)
	}

	return lb, nil
}

// ConfigurationDecodeHook is a composite mapstructure decode hook that contains
// hooks for decoding everything in a Configuration object.
var ConfigurationDecodeHook = mapstructure.ComposeDecodeHookFunc(
	mapToFile,
	stringToBytes,
	mapToLoadBalancer,
)

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

	config := &mapstructure.DecoderConfig{
		DecodeHook: ConfigurationDecodeHook,
		Result:     &response,
	}

	decoder, err := mapstructure.NewDecoder(config)

	if err != nil {
		return nil, err
	}

	err = decoder.Decode(r.Body)

	if err != nil {
		return nil, err
	}

	return &response.Configuration, nil
}

// GetResult temporarily contains the response from a Get call.
type GetResult struct {
	launchResult
}
