package rackspace

import "github.com/rackspace/gophercloud"

type ErrNoAuthURL struct{ *gophercloud.InvalidInputError }

func (e *ErrNoAuthURL) Error() string {
	return "Environment variable RS_AUTH_URL or OS_AUTH_URL need to be set."
}

type ErrNoUsername struct{ *gophercloud.InvalidInputError }

func (e *ErrNoUsername) Error() string {
	return "Environment variable RS_USERNAME or OS_USERNAME need to be set."
}

type ErrNoPassword struct{ *gophercloud.InvalidInputError }

func (e *ErrNoPassword) Error() string {
	return "Environment variable RS_API_KEY or RS_PASSWORD needs to be set."
}
