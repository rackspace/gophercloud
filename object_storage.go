package gophercloud

import (
	"net/http"
	"github.com/racker/perigee"
	"strings"
	"log"
)

const (
	MetadataPrefix = "X-Container-Meta-"
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

// openstackContainer provides the backing state required to keep track of a single container in an OpenStack
// environment.
type openstackContainer struct {
	// Name labels the container.
	Name string

	// Provider links the container to an actual provider.
	Provider *openstackObjectStoreProvider

	// customValues provides access to the custom metadata for this container.
	customValues http.Header
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

func (c *openstackContainer) Metadata() (MetadataProvider, error) {
	// As of this writing, we let the openstackContainer structure keep track of its own metadata.
	return c, nil
}

func (c *openstackContainer) cacheHeaders() error {
	osp := c.Provider
	return osp.context.WithReauth(osp.access, func() error {
		if c.customValues == nil {
			url := osp.endpoint + "/" + c.Name
			resp, err := perigee.Request("HEAD", url, perigee.Options{
				CustomClient: osp.context.httpClient,
				MoreHeaders: map[string]string{
					"X-Auth-Token": osp.access.AuthToken(),
				},
				OkCodes: []int{204},
			})
			if err != nil {
				return err
			}

			c.customValues = resp.HttpResponse.Header
			for key, _ := range c.customValues {
				log.Printf(key)
			}
		}
		return nil
	})
}

// See MetadataProvider interface for details.
func (c *openstackContainer) CustomValues() (map[string]string, error) {
	err := c.cacheHeaders()
	if err != nil {
		return nil, err
	}

	res := map[string]string{}
	for name, values := range c.customValues {
		if strings.HasPrefix(name, MetadataPrefix) {
			res[name] = values[0]
		}
	}
	return res, nil
}

// See MetadataProvider interface for details.
func (c *openstackContainer) CustomValue(key string) (string, error) {
	err := c.cacheHeaders()
	if err != nil {
		return "", err
	}
	value := c.customValues[MetadataPrefix + key]
	if len(value) > 0 {
		return value[0], nil
	}
	return "", nil
}

// See MetadataProvider interface for details.
func (c *openstackContainer) SetCustomValue(key, value string) error {
	osp := c.Provider
	err := osp.context.WithReauth(osp.access, func() error {
		url := osp.endpoint + "/" + c.Name
		_, err := perigee.Request("POST", url, perigee.Options{
			CustomClient: osp.context.httpClient,
			MoreHeaders: map[string]string{
				"X-Auth-Token": osp.access.AuthToken(),
				MetadataPrefix + key: value,
			},
			OkCodes: []int{204},
		})
		return err
	})
	
	// Flush our values cache to make sure our next attempt at getting values always gets the right data.
	if err == nil {
		c.customValues = nil
	}

	return err
}
