package launch

import "errors"

// Validation errors returned by update operations.
var (
	ErrUnknownType      = errors.New("Unknown launch configuration type.")
	ErrNoName           = errors.New("Server name cannot be empty.")
	ErrNoFlavor         = errors.New("Server flavor cannot be empty.")
	ErrNoImage          = errors.New("Server image cannot be empty.")
	ErrUnknownLBType    = errors.New("Unknown load balancer type.")
	ErrNoLoadBalancerID = errors.New("Load balancer ID cannot be empty.")
	ErrNoPort           = errors.New("Cloud load balancer port cannot be zero.")
	ErrNoNetworkID      = errors.New("Network UUID cannot be empty.")
	ErrBadDrainTimeout  = errors.New("Draining timeout must be in range: (30, 3600)")
)
