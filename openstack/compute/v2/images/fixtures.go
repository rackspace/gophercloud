// +build fixtures

package images

import (
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

var (
	testMetadataMap    = map[string]string{"foo": "bar"}
	testMetadataOpts   = MetadatumOpts{"foo": "bar"}
	testMetadataString = `{"metadata": {"foo": "bar"}}`

	testMetadataChangeOpts   = MetadataOpts{"foo": "baz"}
	testMetadataChangeString = `{"metadata": {"foo": "baz"}}`
	testMetadataResetString  = testMetadataChangeString
	testMetadataResetMap     = map[string]string{"foo": "baz"}
	testMetadataUpdateString = `{"metadata": {"foo": "baz"}}`
	testMetadataUpdateMap    = map[string]string{"foo": "baz"}

	testMetadatumMap    = map[string]string{"foo": "bar"}
	testMetadatumOpts   = MetadatumOpts{"foo": "bar"}
	testMetadatumString = `{"meta": {"foo": "bar"}}`

	testMetadatumChangeOpts   = MetadatumOpts{"foo": "bar"}
	testMetadatumChangeString = `{"meta": {"foo": "bar"}}`
	testMetadatumCreateMap    = map[string]string{"foo": "bar"}
	testMetadatumCreateString = testMetadatumChangeString
)

// HandleMetadataGetSuccessfully sets up the test server to respond to a metadata Get request.
func HandleMetadataGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1234asdf/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testMetadataString))
	})
}

// HandleMetadataResetSuccessfully sets up the test server to respond to a metadata Create request.
func HandleMetadataResetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1234asdf/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, testMetadataResetString)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(testMetadataResetString))
	})
}

// HandleMetadataUpdateSuccessfully sets up the test server to respond to a metadata Update request.
func HandleMetadataUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1234asdf/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, testMetadataResetString)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(testMetadataUpdateString))
	})
}

// HandleMetadatumGetSuccessfully sets up the test server to respond to a metadatum Get request.
func HandleMetadatumGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1234asdf/metadata/foo", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(testMetadatumString))
	})
}

// HandleMetadatumCreateSuccessfully sets up the test server to respond to a metadatum Create request.
func HandleMetadatumCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1234asdf/metadata/foo", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, testMetadatumChangeString)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(testMetadatumChangeString))
	})
}

// HandleMetadatumDeleteSuccessfully sets up the test server to respond to a metadatum Delete request.
func HandleMetadatumDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1234asdf/metadata/foo", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}
