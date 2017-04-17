package images

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TestCreateImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Location", "http://localhost:9292/v2/images/b2173dd3-7ad6-4362-baa6-a68bce3565ca")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, `
{
    "status": "queued",
    "name": "Ubuntu 12.10",
    "tags": [
        "ubuntu",
        "quantal"
    ],
    "container_format": "bare",
    "created_at": "2014-11-11T20:47:55Z",
    "disk_format": "qcow2",
    "updated_at": "2014-11-11T20:47:55Z",
    "visibility": "private",
    "self": "/v2/images/b2173dd3-7ad6-4362-baa6-a68bce3565ca",
    "min_disk": 0,
    "protected": false,
    "id": "b2173dd3-7ad6-4362-baa6-a68bce3565ca",
    "file": "/v2/images/b2173dd3-7ad6-4362-baa6-a68bce3565ca/file",
    "owner": "b4eedccc6fb74fa8a7ad6b08382b852b",
    "min_ram": 0,
    "schema": "/v2/schemas/image",
    "size": null,
    "checksum": null,
    "virtual_size": null
}
		`)
	})

	actual, err := Create(fake.ServiceClient(), CreateOpts{
		Name:            "Ubuntu 12.10",
		Tags:            []string{"ubuntu", "quantal"},
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
		Visibility:      "private",
		MinDisk:         0,
		Protected:       false,
		MinRam:          0,
	})
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	if actual == "" {
		t.Errorf("Expected an image id but none found")
	}
}

func TestUploadImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/images/12345/file", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	file, err := ioutil.TempFile(os.TempDir(), "temp")

	if err != nil {
		panic(err)
	}

	tempFilePath, err := filepath.Abs(filepath.Dir(file.Name()))

	if err != nil {
		panic(err)
	}

	err = Upload(fake.ServiceClient(), "12345", tempFilePath)
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	defer os.Remove(file.Name())
}
