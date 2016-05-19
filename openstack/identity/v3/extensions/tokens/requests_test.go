package tokens

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/identity/v3/extensions"
	"github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
	"github.com/rackspace/gophercloud/testhelper"
)

// authTokenPost verifies that providing certain AuthOptions and Scope results in an expected JSON structure.
func authTokenPost(t *testing.T, options extensions.AuthOptions, scope *Scope, requestJSON string) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	client := gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{
			TokenID: "12345abcdef",
		},
		Endpoint: testhelper.Endpoint(),
	}

	testhelper.Mux.HandleFunc("/auth/tokens", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "POST")
		testhelper.TestHeader(t, r, "Content-Type", "application/json")
		testhelper.TestHeader(t, r, "Accept", "application/json")
		testhelper.TestJSONRequest(t, r, requestJSON)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
			"token": {
				"expires_at": "2014-10-02T13:45:00.000000Z"
			}
		}`)
	})

	_, err := Create(&client, options, scope).Extract()
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}
}

func authTokenPostErr(t *testing.T, options extensions.AuthOptions, scope *Scope, includeToken bool, expectedErr error) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	client := gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{},
		Endpoint:       testhelper.Endpoint(),
	}
	if includeToken {
		client.TokenID = "abcdef123456"
	}

	_, err := Create(&client, options, scope).Extract()
	if err == nil {
		t.Errorf("Create did NOT return an error")
	}
	if err != expectedErr {
		t.Errorf("Create returned an unexpected error: wanted %v, got %v", expectedErr, err)
	}
}

func TestTrustIDTokenID(t *testing.T) {
	options := extensions.AuthOptions{&gophercloud.AuthOptions{TokenID: "old_trustee"}, "123456"}
	scope := &Scope{TrustID: "123456"}
	authTokenPost(t, options, scope, `
		{
		  "auth": {
		    "identity": {
			      "methods": [
		        	"token"
			      ],
				"token": {
			        	"id": "12345abcdef"
			      }
		    },
		    "scope": {
			      "OS-TRUST:trust": {
			        "id": "123456"
			      }
			    }
			}
		}

	`)
}

func TestFailurePassword(t *testing.T) {
	options := extensions.AuthOptions{&gophercloud.AuthOptions{TokenID: "fakeidnopass"}, "123456"}
	//Service Client must have tokenId or password,
	//setting include tokenId to false
	scope := &Scope{TrustID: "notenough"}
	authTokenPostErr(t, options, scope, false, tokens.ErrMissingPassword)
}
