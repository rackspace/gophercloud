package images

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"io"
	"log"
	"os"
	"strings"
)

type CreateOpts struct {
	Name            string                 `json:"name,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
	ID              string                 `json:"id,omitempty"`
	Visibility      string                 `json:"visibility,omitempty"`
	ContainerFormat string                 `json:"container_format,omitempty"`
	DiskFormat      string                 `json:"disk_format,omitempty"`
	MinDisk         int                    `json:"min_disk,omitempty"`
	MinRam          int                    `json:"min_ram,omitempty"`
	Protected       bool                   `json:"protected,omitempty"`
	Properties      map[string]interface{} `json:"properties,omitempty"`
}

// Add adds a new image with the specified meta-data.
func Create(client *gophercloud.ServiceClient, opts CreateOpts) (string, error) {

	resp, err := client.Request("POST", addURL(client), gophercloud.RequestOpts{
		OkCodes:  []int{201},
		JSONBody: opts,
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	location := resp.Header.Get("location")

	// Return the last element of the location which is the image id
	locationArr := strings.Split(location, "/")
	return locationArr[len(locationArr)-1], err
}

func streamFile(readFrom *os.File, readFromPath string, writePipe *io.PipeWriter, formLabel string, ppErr **error) {

	// Assure the file closes when exiting this function. Note that the
	// caller should not defer this close since this function likely runs
	// asynchronously.
	defer readFrom.Close()

	// Assure the write side of the pipe closes when exiting this function.
	defer writePipe.Close()

	// copy from the file to stream into the multipart.
	_, err := io.Copy(writePipe, readFrom)
	if err != nil {
		*ppErr = &err
		return
	}

	*ppErr = nil
}

func Upload(client *gophercloud.ServiceClient, imageId string, imagePath string) error {
	image, err := os.Stat(imagePath)
	if err != nil {
		return err
	}
	imageSize := fmt.Sprintf("%d", image.Size())

	headers := map[string]string{
		"Content-Length": imageSize,
	}

	for k, v := range client.AuthenticatedHeaders() {
		headers[k] = v
	}

	inImage, err := os.Open(imagePath)
	if err != nil {
		return err
	}

	body, writer := io.Pipe()

	var streamErr *error
	go streamFile(inImage, imagePath, writer, "file", &streamErr)

	_, err = client.Request("PUT", uploadUrl(client, imageId), gophercloud.RequestOpts{
		MoreHeaders:      headers,
		OkCodes:          []int{204},
		StreamingRawBody: body,
	})

	return err
}
