package images

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// GetResult temporarily stores a Get response.
type GetResult struct {
	gophercloud.Result
}

// DeleteResult represents the result of an image.Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Extract interprets a GetResult as an Image.
func (gr GetResult) Extract() (*Image, error) {
	if gr.Err != nil {
		return nil, gr.Err
	}

	var decoded struct {
		Image Image `mapstructure:"image"`
	}

	err := mapstructure.Decode(gr.Body, &decoded)
	return &decoded.Image, err
}

// Image is used for JSON (un)marshalling.
// It provides a description of an OS image.
type Image struct {
	// ID contains the image's unique identifier.
	ID string

	Created string

	// MinDisk and MinRAM specify the minimum resources a server must provide to be able to install the image.
	MinDisk int
	MinRAM  int

	// Name provides a human-readable moniker for the OS image.
	Name string

	// The Progress and Status fields indicate image-creation status.
	// Any usable image will have 100% progress.
	Progress int
	Status   string

	Updated string

	Metadata map[string]string
}

// ImagePage contains a single page of results from a List operation.
// Use ExtractImages to convert it into a slice of usable structs.
type ImagePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no Image results.
func (page ImagePage) IsEmpty() (bool, error) {
	images, err := ExtractImages(page)
	if err != nil {
		return true, err
	}
	return len(images) == 0, nil
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page ImagePage) NextPageURL() (string, error) {
	type resp struct {
		Links []gophercloud.Link `mapstructure:"images_links"`
	}

	var r resp
	err := mapstructure.Decode(page.Body, &r)
	if err != nil {
		return "", err
	}

	return gophercloud.ExtractNextURL(r.Links)
}

// ExtractImages converts a page of List results into a slice of usable Image structs.
func ExtractImages(page pagination.Page) ([]Image, error) {
	casted := page.(ImagePage).Body
	var results struct {
		Images []Image `mapstructure:"images"`
	}

	err := mapstructure.Decode(casted, &results)
	return results.Images, err
}

// MetadataResult contains the result of a call for (potentially) multiple key-value pairs.
type MetadataResult struct {
	gophercloud.Result
}

// GetMetadataResult temporarily contains the response from a metadata Get call.
type GetMetadataResult struct {
	MetadataResult
}

// ResetMetadataResult temporarily contains the response from a metadata Reset call.
type ResetMetadataResult struct {
	MetadataResult
}

// UpdateMetadataResult temporarily contains the response from a metadata Update call.
type UpdateMetadataResult struct {
	MetadataResult
}

// MetadatumResult contains the result of a call for individual a single key-value pair.
type MetadatumResult struct {
	gophercloud.Result
}

// GetMetadatumResult temporarily contains the response from a metadatum Get call.
type GetMetadatumResult struct {
	MetadatumResult
}

// CreateMetadatumResult temporarily contains the response from a metadatum Create call.
type CreateMetadatumResult struct {
	MetadatumResult
}

// DeleteMetadatumResult temporarily contains the response from a metadatum Delete call.
type DeleteMetadatumResult struct {
	gophercloud.ErrResult
}

// Extract interprets any MetadataResult as a Metadata, if possible.
func (r MetadataResult) Extract() (map[string]string, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Metadata map[string]string `mapstructure:"metadata"`
	}

	err := mapstructure.Decode(r.Body, &response)
	return response.Metadata, err
}

// Extract interprets any MetadatumResult as a Metadatum, if possible.
func (r MetadatumResult) Extract() (map[string]string, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Metadatum map[string]string `mapstructure:"meta"`
	}

	err := mapstructure.Decode(r.Body, &response)
	return response.Metadatum, err
}
