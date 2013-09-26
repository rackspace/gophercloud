package gophercloud

import (
	"github.com/racker/perigee"
)

// The openstackObjectStorageProvider structure provides the implementation for generic OpenStack-compatible
// object storage interfaces.
type openstackObjectStoreProvider struct {
	// endpoint refers to the provider's API endpoint base URL.  This will be used to construct
	// and issue queries.
	endpoint string

	// Test context (if any) in which to issue requests.
	context *Context

	// access associates this API provider with a set of credentials,
	// which may be automatically renewed if they near expiration.
	access AccessProvider
}

type openstackContainer struct {
	// Name labels the container.
	Name string

	// Provider links the container to an actual provider.
	Provider ObjectStoreProvider
}

func (osp *openstackObjectStoreProvider) CreateContainer(name string) (Container, error) {
	var container Container

	err := osp.context.WithReauth(osp.access, func() error {
		url := osp.endpoint + "/" + name
		err := perigee.Put(url, perigee.Options{
			CustomClient: osp.context.httpClient,
			MoreHeaders: map[string]string{
				"X-Auth-Token": osp.access.AuthToken(),
			},
			OkCodes: []int{201},
		})
		if err == nil {
			container = &openstackContainer{
				Name: name,
				Provider: osp,
			}
		}
		return err
	})
	return container, err
}

func (osp *openstackObjectStoreProvider) DeleteContainer(name string) error {
	err := osp.context.WithReauth(osp.access, func() error {
		url := osp.endpoint + "/" + name
		return perigee.Delete(url, perigee.Options{
			CustomClient: osp.context.httpClient,
			MoreHeaders: map[string]string{
				"X-Auth-Token": osp.access.AuthToken(),
			},
			OkCodes: []int{204},
		})
	})
	return err
}

func (c *openstackContainer) Delete() error {
	return c.Provider.DeleteContainer(c.Name)
}
