package softlayer

import (
	"encoding/json"
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/utils"
	"net/http"
	"net/url"
)

const (
	//	v20 = "v2.0"
	//	v30 = "v3.0"
	v10 = "v1.0"
)

// AuthenticatedClient logs in to an OpenStack cloud found at the identity endpoint specified by options, acquires a token, and
// returns a Client instance that's ready to operate.
// It first queries the root identity endpoint to determine which versions of the identity service are supported, then chooses
// the most recent identity service available to proceed.
func AuthenticatedClient(options gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	client, err := NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewClient prepares an unauthenticated ProviderClient instance.
// Most users will probably prefer using the AuthenticatedClient function instead.
// This is useful if you wish to explicitly control the version of the identity service that's used for authentication explicitly,
// for example.
func NewClient(endpoint string) (*gophercloud.ProviderClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	hadPath := u.Path != ""
	u.Path, u.RawQuery, u.Fragment = "", "", ""
	base := u.String()

	endpoint = gophercloud.NormalizeURL(endpoint)
	base = gophercloud.NormalizeURL(base)

	if hadPath {
		return &gophercloud.ProviderClient{
			IdentityBase:     base,
			IdentityEndpoint: endpoint,
		}, nil
	}

	return &gophercloud.ProviderClient{
		IdentityBase:     base,
		IdentityEndpoint: "",
	}, nil
}

// AuthenticatedClient logs in to an OpenStack cloud found at the identity endpoint specified by options, acquires a token, and
// returns a Client instance that's ready to operate.
// It first queries the root identity endpoint to determine which versions of the identity service are supported, then chooses
// the most recent identity service available to proceed.
//func AuthenticatedClient(options gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
//	client, err := NewClient(options.IdentityEndpoint)
//	if err != nil {
//		return nil, err
//	}
//
//	err = Authenticate(client, options)
//	if err != nil {
//		return nil, err
//	}
//	return client, nil
//}

// Authenticate or re-authenticate against the most recent identity service supported at the provided endpoint.
func Authenticate(client *gophercloud.ProviderClient, options gophercloud.AuthOptions) error {
	versions := []*utils.Version{
		&utils.Version{ID: v10, Priority: 10, Suffix: "/v1.0/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return err
	}

	switch chosen.ID {
	case v10:
		return v1auth(client, endpoint, options)
	default:
		// The switch statement must be out of date from the versions list.
		return fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
//func AuthenticateV2(client *gophercloud.ProviderClient, options gophercloud.AuthOptions) error {
//	return v2auth(client, "", options)
//}
type SoftLayerObjectStoreError struct {
	Status_Code int
	Message     string
}

func (e SoftLayerObjectStoreError) Error() string { return e.Message }

// Performs an HTTP request.
func DoRequest(req *http.Request, data interface{}) (resp *http.Response, err error) {
	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, &SoftLayerObjectStoreError{resp.StatusCode, resp.Status}
	}

	if data != nil {
		err = decoder.Decode(data)
		if err != nil {
			return
		}
	}
	return
}

func v1auth(client *gophercloud.ProviderClient, endpoint string, options gophercloud.AuthOptions) error {
	v2Client := NewIdentityV1(client)
	if endpoint != "" {
		v2Client.Endpoint = endpoint
	}

	//TODO: send to SL and get header token , set it to token in
	//	http.NewRequest(method, client.Base_Url+"/"+endpoint, bytes.NewBuffer(body))
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Auth-User", options.Username)
	req.Header.Set("X-Auth-Key", options.APIKey)

	resp, err := DoRequest(req, nil)
	if err != nil {
		return err
	}
	//	client.Username = username
	//	client.Api_Key = api_key
	//	client.Auth_Token = resp.Header["X-Auth-Token"][0]
	//	client.Authenticated = true

	client.TokenID = resp.Header["X-Auth-Token"][0]
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		//TODO: fix me
		// https://dal05.objectstorage.softlayer.net/v1/AUTH_df0de35c-d00a-40aa-b697-2b7f1b9331a6
		//		return V2EndpointURL(catalog, opts)
		return resp.Header["X-Storage-Url"][0], nil
	}

	return nil
}

// NewIdentityV2 creates a ServiceClient that may be used to interact with the v2 identity service.
func NewIdentityV1(client *gophercloud.ProviderClient) *gophercloud.ServiceClient {
	v2Endpoint := client.IdentityBase + "v1.0/"

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       v2Endpoint,
	}
}

// NewObjectStorageV1 creates a ServiceClient that may be used with the v1 object storage package.
func NewObjectStorageV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	eo.ApplyDefaults("object-store")
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return nil, err
	}
	return &gophercloud.ServiceClient{ProviderClient: client, Endpoint: url}, nil
}
