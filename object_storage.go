package gophercloud

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/racker/perigee"
	"reflect"
	"strconv"
	"strings"
)

const (
	ContainerMetadataPrefix = "x-container-meta-"
)

// containerMetaName takes an unadorned custom metadata key and formats it suitably for map
// look-up.
func containerMetaName(s string) string {
	return strings.ToLower(ContainerMetadataPrefix + s)
}

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

	// customMetadata provides access to the custom metadata for this container.
	customMetadata *cimap
}

// openstackContainerInfo holds the information describing a single OpenStack container.
type openstackContainerInfo struct {
	// Bytes is the the size of the container.
	Bytes int
	// Count is the number of objects in the container.
	Count int
	// Name is the label for the container.
	Name string
}

// openstackObjectInfo holds the information describing a single OpenStack object.
type openstackObjectInfo struct {
	Name          string
	Hash          string
	Bytes         int
	Content_type  string
	Last_modified string
}

// See ObjectStoreProvider interface for details.
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
				Name:     name,
				Provider: osp,
			}
		}
		return err
	})
	return container, err
}

// See ObjectStoreProvider interface for details.
func (osp *openstackObjectStoreProvider) ListContainers(listOpts ListOpts) ([]ContainerInfo, error) {
	returnFull := listOpts.Full
	if returnFull {
		var osci []openstackContainerInfo
		err := osp.context.WithReauth(osp.access, func() error {
			url := osp.endpoint
			_, err := perigee.Request("GET", url, perigee.Options{
				CustomClient: osp.context.httpClient,
				Results:      &osci,
				MoreHeaders: map[string]string{
					"X-Auth-Token": osp.access.AuthToken(),
				},
			})
			return err
		})
		containersInfo := make([]ContainerInfo, len(osci))
		for i, val := range osci {
			containersInfo[i] = val
		}
		return containersInfo, err
	} else {
		response, err := osp.context.ResponseWithReauth(osp.access, func() (*perigee.Response, error) {
			url := osp.endpoint
			return perigee.Request("GET", url, perigee.Options{
				CustomClient: osp.context.httpClient,
				Results:      true,
				MoreHeaders: map[string]string{
					"X-Auth-Token": osp.access.AuthToken(),
				},
				Accept: "text/plain",
			})
		})
		rawResult := string(response.JsonResult)
		containerNames := strings.Split(rawResult[:len(rawResult)-1], "\n")
		containersInfo := make([]ContainerInfo, len(containerNames))
		for i, containerName := range containerNames {
			containersInfo[i] = openstackContainerInfo{
				Name: containerName,
			}
		}
		return containersInfo, err
	}
}

// See ObjectStoreProvider interface for details.
func (osp *openstackObjectStoreProvider) GetContainer(name string) Container {
	return &openstackContainer{
		Name:     name,
		Provider: osp,
	}
}

// See ObjectStoreProvider interface for details
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

// See Container interface for details
func (c *openstackContainer) ListObjects(listOpts ListOptions) ([]ObjectInfo, error) {
	osListOpts, ok := listOpts.(OpenstackListOpts)
	if !ok {
		return nil, errors.New("Error casting from interface ListOptions to structure OpenstackListOpts.")
	}
	queryString := "?"
	l := reflect.ValueOf(&osListOpts).Elem()
	typeOfL := l.Type()
	for i := 0; i < l.NumField(); i++ {
		f := l.Field(i)
		fName := typeOfL.Field(i).Name
		fValue := f.Interface()
		switch f.Kind() {
		case reflect.String:
			if fValue.(string) != "" {
				queryString += fName + "=" + fValue.(string)
			}
		case reflect.Int:
			if fValue.(int) != 0 {
				queryString += fName + "=" + strconv.Itoa(fValue.(int))
			}
		case reflect.Slice:
			if fValue.([]byte) != nil {
				queryString += fName + "=" + string(fValue.([]byte))
			}
		}
	}
	if queryString != "?" {
		queryString = strings.ToLower(queryString)
	}
	osp := c.Provider
	url := fmt.Sprintf("%s/%s%s", osp.endpoint, c.Name, queryString)
	returnFull := osListOpts.Full
	if returnFull {
		var osoi []openstackObjectInfo
		err := osp.context.WithReauth(osp.access, func() error {
			_, err := perigee.Request("GET", url, perigee.Options{
				CustomClient: osp.context.httpClient,
				Results:      &osoi,
				MoreHeaders: map[string]string{
					"X-Auth-Token": osp.access.AuthToken(),
				},
				OkCodes: []int{200, 204},
			})

			return err
		})
		objectsInfo := make([]ObjectInfo, len(osoi))
		for i, val := range osoi {
			objectsInfo[i] = val
		}

		return objectsInfo, err
	} else {
		response, err := osp.context.ResponseWithReauth(osp.access, func() (*perigee.Response, error) {
			return perigee.Request("GET", url, perigee.Options{
				CustomClient: osp.context.httpClient,
				Results:      true,
				MoreHeaders: map[string]string{
					"X-Auth-Token": osp.access.AuthToken(),
				},
				OkCodes: []int{200},
				Accept:  "text/plain",
			})
		})
		rawResult := string(response.JsonResult)
		objectNames := strings.Split(rawResult[:len(rawResult)-1], "\n")
		objectsInfo := make([]ObjectInfo, len(objectNames))
		for i, objectName := range objectNames {
			objectsInfo[i] = openstackObjectInfo{
				Name: objectName,
			}
		}

		return objectsInfo, err
	}
}

func (c *openstackContainer) Delete() error {
	return c.Provider.DeleteContainer(c.Name)
}

func (c *openstackContainer) Metadata() (MetadataProvider, error) {
	// As of this writing, we let the openstackContainer structure keep track of its own metadata.
	return c, nil
}

// cacheHeaders() takes no action if custom metadata headers have already been retrieved.
// Otherwise, the container resource is queried for its current set of custom headers.
func (c *openstackContainer) cacheHeaders() error {
	osp := c.Provider
	return osp.context.WithReauth(osp.access, func() error {
		if c.customMetadata == nil {
			// Grab the set of headers attached to this container.
			// These headers will be keyed off of mixed-case strings.
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

			// To ensure case insensitivity when looking up keys,
			// transcribe our headers such that all the keys used to
			// index them are lower-case.
			headers := resp.HttpResponse.Header
			loweredHeaders := make(map[string]string)
			for key, values := range headers {
				key = strings.ToLower(key)
				if strings.HasPrefix(key, containerMetaName("")) {
					loweredHeaders[key[len(ContainerMetadataPrefix):]] = values[0]
				}
			}
			c.customMetadata = &cimap{m: loweredHeaders}
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
	return c.customMetadata.rawMap(), nil
}

// See MetadataProvider interface for details.
func (c *openstackContainer) CustomValue(key string) (string, error) {
	err := c.cacheHeaders()
	if err != nil {
		return "", err
	}
	key = strings.ToLower(key)
	value, _ := c.customMetadata.get(key)
	if len(value) > 0 {
		return value, nil
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
				"X-Auth-Token":         osp.access.AuthToken(),
				containerMetaName(key): value,
			},
			OkCodes: []int{204},
		})
		return err
	})

	// Flush our values cache to make sure our next attempt at getting values always gets the right data.
	if err == nil {
		c.customMetadata = nil
	}

	return err
}

// BasicObjectUploader returns a pointer to a BasicUploader object with an empty buffer
func (c *openstackContainer) BasicObjectUploader() *BasicUploader {
	return &BasicUploader{bytes.NewBuffer(make([]byte, 0))}
}

// See Container interface for details.
func (c *openstackContainer) DeleteObject(name string) error {
	osp := c.Provider
	return osp.context.WithReauth(osp.access, func() error {
		url := fmt.Sprintf("%s/%s/%s", osp.endpoint, c.Name, name)
		_, err := perigee.Request("DELETE", url, perigee.Options{
			CustomClient: osp.context.httpClient,
			MoreHeaders: map[string]string{
				"X-Auth-Token": osp.access.AuthToken(),
			},
			OkCodes: []int{204},
		})
		return err
	})
}

// See ContainerInfo interface for details
func (ci openstackContainerInfo) Label() string {
	return ci.Name
}

// See ContainerInfo interface for details
func (ci openstackContainerInfo) ObjCount() int {
	return ci.Count
}

// See ContainerInfo interface for details
func (ci openstackContainerInfo) Size() int {
	return ci.Bytes
}

func (oi openstackObjectInfo) GetName() string {
	return oi.Name
}

func (oi openstackObjectInfo) GetHash() string {
	return oi.Hash
}

func (oi openstackObjectInfo) GetSize() int {
	return oi.Bytes
}

func (oi openstackObjectInfo) GetContentType() string {
	return oi.Content_type
}

func (oi openstackObjectInfo) GetLastModified() string {
	return oi.Last_modified
}

// Commit attempts to upload the object data to the endpoint.
func (bu *BasicUploader) Commit(objOpts ObjectOpts) error {
	c := objOpts.Container.(*openstackContainer)
	osp := c.Provider
	err := osp.context.WithReauth(osp.access, func() error {
		url := fmt.Sprintf("%s/%s/%s", osp.endpoint, c.Name, objOpts.Name)
		moreHeaders := map[string]string{
			"X-Auth-Token": osp.access.AuthToken(),
		}

		reqBody := make([]byte, bu.Len())
		_, err := bu.Read(reqBody)
		if err != nil {
			return err
		}

		_, err = perigee.Request("PUT", url, perigee.Options{
			CustomClient: osp.context.httpClient,
			ReqBody:      reqBody,
			MoreHeaders:  moreHeaders,
			DumpReqJson:  true,
			OkCodes:      []int{201},
		})

		return err
	})

	return err
}

// *BasicUploader.WriteAt writes a slice of bytes (p) at a particular offset (off).
// It is used to seek in the buffer. Seeking beyond the bounds of the already-written
// object results in an error.
func (bu *BasicUploader) WriteAt(p []byte, off int64) (int, error) {
	if off > int64(bu.Len()) || off+int64(len(p)) > int64(bu.Len()) {
		return 0, errors.New("Slice bounds out of range.")
	}
	curBytes := bu.Bytes()
	newData := bytes.Replace(curBytes, curBytes[off:off+int64(len(p))], p, 1)
	bu.Reset()
	_, err := bu.Write(newData)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// *BasicUploader.Close 'closes' the BasicUploader by nilling the pointer.
func (bu *BasicUploader) Close() error {
	bu = nil
	return nil
}

// BasicUploader
type BasicUploader struct {
	*bytes.Buffer
}

// ObjectOpts is a structure containing relevant parameters when creating an uploader or downloader.
type ObjectOpts struct {
	Length    int
	Name      string
	Offset    int
	Container Container
}

// ListOpts is a structure containing relevant parameters when requesting a list of containers.
type ListOpts struct {
	Full      bool
	Limit     int
	Marker    string
	EndMarker string
}

// OpenstackListOpts is a structure containing relevant parameters when listing items in OpenStack.
type OpenstackListOpts struct {
	Full      bool
	Limit     int
	Marker    string
	EndMarker string
	Prefix    string
	Format    string
	Delimiter []byte
	Path      string
}

// See ListOptions interface for details.
func (lo OpenstackListOpts) GetFull() bool {
	return lo.Full
}

// See ListOptions interface for details.
func (lo OpenstackListOpts) GetLimit() int {
	return lo.Limit
}

// See ListOptions interface for details.
func (lo OpenstackListOpts) GetMarker() string {
	return lo.Marker
}

// See ListOptions interface for details.
func (lo OpenstackListOpts) GetEndMarker() string {
	return lo.EndMarker
}
