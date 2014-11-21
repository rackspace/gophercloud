package floatingips

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud/pagination"
)

// FloatingIP represents a floating IP address, which is a public
// IP address that can be associated with a server instance.
type FloatingIP struct {
	// ID is the unique identitfier of the floating IP address.
	ID string `mapstructure:"id"`
	// IP is the IP address of the floating IP address.
	IP string `mapstructure:"ip"`
	// InstanceID is the ID of the server instance with which the floating IP
	// address is associated, if any.
	InstanceID string `mapstructure:"instance_id"`
	// Pool is the name of the pool to which the floating IP address belongs.
	Pool string `mapstructure:"pool"`
	// FixedIP is the fixed IP address of the server instance to which the floating
	// IP address is associated, if any.
	FixedIP string `mapstructure:"fixed_ip"`
}

// FloatingIPPage stores a single page of FloatingIP results from a List call.
type FloatingIPPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a FloatingIPPage is empty.
func (page FloatingIPPage) IsEmpty() (bool, error) {
	ks, err := ExtractFloatingIPs(page)
	return len(ks) == 0, err
}

// ExtractFloatingIPs interprets a page of results as a slice of FloatingIPs.
func ExtractFloatingIPs(page pagination.Page) ([]FloatingIP, error) {
	var resp struct {
		FloatingIPs []FloatingIP `mapstructure:"floating_ips"`
	}

	err := mapstructure.Decode(page.(FloatingIPPage).Body, &resp)
	return resp.FloatingIPs, err
}
