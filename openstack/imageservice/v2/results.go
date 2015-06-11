package v2

import (
	"errors"
	"fmt"
	
	"github.com/rackspace/gophercloud"
	//"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

// does not include the literal image data; just metadata.
// returned by listing images, and by fetching a specific image.
type Image struct {
	Id string
	
	Name string
	
	Status ImageStatus
	
	Tags []string
	
	ContainerFormat string `mapstructure:"container_format"`
	DiskFormat string `mapstructure:"disk_format"`
	
	MinDiskGigabytes int `mapstructure:"min_disk"`
	MinRamMegabytes int `mapstructure:"min_ram"`
	
	Owner string `mapstructure:"owner"`
	
	Protected bool `mapstructure:"protected"`
	Visibility ImageVisibility `mapstructure:"visibility"`

	Checksum string `mapstructure:"checksum"`
	SizeBytes int `mapstructure:"size"`
	
	Metadata map[string]string `mapstructure:"metadata"`
	Properties map[string]string `mapstructure:"properties"`
}

// implements pagination.Page<Image>, pagination.MarkerPage
// DOESN'T implement Page. Why? How does openstack/compute/v2
// type ImagePage struct {
// 	pagination.MarkerPageBase  // pagination.MarkerPageBase<Image>
// }

type CreateResult struct {
	gophercloud.ErrResult
}

func asString(any interface{}) (string, error) {
	if str, ok := any.(string); ok {
		return str, nil
	} else {
		return "", errors.New(fmt.Sprintf("expected string value, but found: %#v", any))
	}
}

func extractStringAtKey(m map[string]interface{}, k string) (string, error) {
	if any, ok := m[k]; ok {
		return asString(any)
	} else {
		return "", errors.New(fmt.Sprintf("expected key \"%s\" in map, but this key is not present", k))
	}
}

func extractStringSliceAtKey(m map[string]interface{}, k string) ([]string, error) {
	if any, ok := m[k]; ok {
		if slice, ok := any.([]interface{}); ok {
			res := make([]string, len(slice))
			for k, v := range slice {
				var err error
				if res[k], err = asString(v); err != nil {
					return nil, err
				}
			}
			return res, nil
		} else {
			return nil, errors.New(fmt.Sprintf("expected slice as \"%s\" value, but found: %#v", k, any))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("expected key \"%s\" in map, but this key is not present", k))
	}
}

func stringToImageStatus(s string) (ImageStatus, error) {
	if s == "queued" {
		return ImageStatusQueued, nil
	} else if s == "active" {
		return ImageStatusActive, nil
	} else {
		return "", errors.New(fmt.Sprintf("expected \"active\" as image status, but found: \"%s\"", s))
	}
}

func extractImageStatusAtKey(m map[string]interface{}, k string) (ImageStatus, error) {
	if any, ok := m[k]; ok {
		if str, ok := any.(string); ok {
			return stringToImageStatus(str)
		} else {
			return "", errors.New(fmt.Sprintf("expected string as \"%s\" value, but found: %#v", k, any))
		}
	} else {
		return "", errors.New(fmt.Sprintf("expected key \"%s\" in map, but this key is not present", k))
	}
}

func extractImage(res gophercloud.ErrResult) (*Image, error) {
	if res.Err != nil {
		return nil, res.Err
	}

	body, ok := res.Body.(map[string]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("expected map as result body, but found: %#v", res.Body))
	}
	
	var image Image

	var err error
	
	if image.Id, err = extractStringAtKey(body, "id"); err != nil {
		return nil, err
	}

	if image.Name, err = extractStringAtKey(body, "name"); err != nil {
		return nil, err
	}

	if image.Status, err = extractImageStatusAtKey(body, "status"); err != nil {
		return nil, err
	}

	if image.Tags, err = extractStringSliceAtKey(body, "tags"); err != nil {
		return nil, err
	}

	return &image, nil
}

// The response to `POST /images` follows the same schema as `GET /images/:id`.
func extractImageOld(res gophercloud.ErrResult) (*Image, error) {
	if res.Err != nil {
		return nil, res.Err
	}

	var image Image

	err := mapstructure.Decode(res.Body, &image)
	
	return &image, err
}

func (c CreateResult) Extract() (*Image, error) {
	return extractImage(c.ErrResult)
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	gophercloud.ErrResult
}

func (c GetResult) Extract() (*Image, error) {
	return extractImage(c.ErrResult)
}

type UpdateResult struct {
	gophercloud.ErrResult
}