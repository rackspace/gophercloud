package images

import (
	"errors"
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToImageListQuery() (string, error)
}

// ListOpts contain options for limiting the number of Images returned from a call to ListDetail.
type ListOpts struct {
	// When the image last changed status (in date-time format).
	ChangesSince string `q:"changes-since"`
	// The number of Images to return.
	Limit int `q:"limit"`
	// UUID of the Image at which to set a marker.
	Marker string `q:"marker"`
	// The name of the Image.
	Name string `q:"name"`
	// The name of the Server (in URL format).
	Server string `q:"server"`
	// The current status of the Image.
	Status string `q:"status"`
	// The value of the type of image (e.g. BASE, SERVER, ALL)
	Type string `q:"type"`
}

// ToImageListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToImageListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// ListDetail enumerates the available images.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailURL(client)
	if opts != nil {
		query, err := opts.ToImageListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	createPage := func(r pagination.PageResult) pagination.Page {
		return ImagePage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, url, createPage)
}

// Get acquires additional detail about a specific image by ID.
// Use GetResult.Extract() to interpret the result as an openstack Image.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getURL(client, id), &result.Body, nil)
	return result
}

// Delete deletes the specified image ID.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var result DeleteResult
	_, result.Err = client.Delete(deleteURL(client, id), nil)
	return result
}

// IDFromName is a convienience function that returns an image's ID given its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	imageCount := 0
	imageID := ""
	if name == "" {
		return "", fmt.Errorf("An image name must be provided.")
	}
	pager := ListDetail(client, &ListOpts{
		Name: name,
	})
	pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := ExtractImages(page)
		if err != nil {
			return false, err
		}

		for _, i := range imageList {
			if i.Name == name {
				imageCount++
				imageID = i.ID
			}
		}
		return true, nil
	})

	switch imageCount {
	case 0:
		return "", fmt.Errorf("Unable to find image: %s", name)
	case 1:
		return imageID, nil
	default:
		return "", fmt.Errorf("Found %d images matching %s", imageCount, name)
	}
}

// Metadata requests all the metadata for the given image ID.
func Metadata(client *gophercloud.ServiceClient, id string) GetMetadataResult {
	var res GetMetadataResult
	_, res.Err = client.Get(metadataURL(client, id), &res.Body, nil)
	return res
}

// MetadataOpts is a map that contains key-value pairs.
type MetadataOpts map[string]string

// ToMetadataResetMap assembles a body for a Reset request based on the contents of a MetadataOpts.
func (opts MetadataOpts) ToMetadataResetMap() (map[string]interface{}, error) {
	return map[string]interface{}{"metadata": opts}, nil
}

// ToMetadataUpdateMap assembles a body for an Update request based on the contents of a MetadataOpts.
func (opts MetadataOpts) ToMetadataUpdateMap() (map[string]interface{}, error) {
	return map[string]interface{}{"metadata": opts}, nil
}

// ResetMetadataOptsBuilder allows extensions to add additional parameters to the
// Reset request.
type ResetMetadataOptsBuilder interface {
	ToMetadataResetMap() (map[string]interface{}, error)
}

// ResetMetadata will create multiple new key-value pairs for the given image ID.
// Note: Using this operation will erase any already-existing metadata and create
// the new metadata provided. To keep any already-existing metadata, use the
// UpdateMetadatas or UpdateMetadata function.
func ResetMetadata(client *gophercloud.ServiceClient, id string, opts ResetMetadataOptsBuilder) ResetMetadataResult {
	var res ResetMetadataResult
	metadata, err := opts.ToMetadataResetMap()
	if err != nil {
		res.Err = err
		return res
	}
	_, res.Err = client.Put(metadataURL(client, id), metadata, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

// UpdateMetadataOptsBuilder allows extensions to add additional parameters to the
// Create request.
type UpdateMetadataOptsBuilder interface {
	ToMetadataUpdateMap() (map[string]interface{}, error)
}

// UpdateMetadata updates (or creates) all the metadata specified by opts for the given image ID.
// This operation does not affect already-existing metadata that is not specified
// by opts.
func UpdateMetadata(client *gophercloud.ServiceClient, id string, opts UpdateMetadataOptsBuilder) UpdateMetadataResult {
	var res UpdateMetadataResult
	metadata, err := opts.ToMetadataUpdateMap()
	if err != nil {
		res.Err = err
		return res
	}
	_, res.Err = client.Post(metadataURL(client, id), metadata, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

// Metadatum requests the key-value pair with the given key for the given image ID.
func Metadatum(client *gophercloud.ServiceClient, id, key string) GetMetadatumResult {
	var res GetMetadatumResult
	_, res.Err = client.Request("GET", metadatumURL(client, id, key), gophercloud.RequestOpts{
		JSONResponse: &res.Body,
	})
	return res
}

// MetadatumOpts is a map of length one that contains a key-value pair.
type MetadatumOpts map[string]interface{}

// ToMetadatumCreateMap assembles a body for a Create request based on the contents of a MetadataumOpts.
func (opts MetadatumOpts) ToMetadatumCreateMap() (map[string]interface{}, string, error) {
	if len(opts) != 1 {
		return nil, "", errors.New("CreateMetadatum operation must have 1 and only 1 key-value pair.")
	}
	metadatum := map[string]interface{}{"meta": opts}
	var key string
	for k := range metadatum["meta"].(MetadatumOpts) {
		key = k
	}
	return metadatum, key, nil
}

// MetadatumOptsBuilder allows extensions to add additional parameters to the
// Create request.
type MetadatumOptsBuilder interface {
	ToMetadatumCreateMap() (map[string]interface{}, string, error)
}

// CreateMetadatum will create or update the key-value pair with the given key for the given image ID.
func CreateMetadatum(client *gophercloud.ServiceClient, id string, opts MetadatumOptsBuilder) CreateMetadatumResult {
	var res CreateMetadatumResult
	metadatum, key, err := opts.ToMetadatumCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Put(metadatumURL(client, id, key), metadatum, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

// DeleteMetadatum will delete the key-value pair with the given key for the given image ID.
func DeleteMetadatum(client *gophercloud.ServiceClient, id, key string) DeleteMetadatumResult {
	var res DeleteMetadatumResult
	_, res.Err = client.Delete(metadatumURL(client, id, key), nil)
	return res
}
