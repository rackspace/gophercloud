package tokens

import "github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
